# Autoparts-Pro

A small, example auto parts inventory and search application. Autoparts-Pro is a Go (Gin) backend API with an Angular frontend, PostgreSQL for storage, Redis for cache/session data, and OpenFGA for authorization.

This draft README focuses on a concise Quick Start, local development notes, configuration, testing, and troubleshooting to help contributors get productive quickly.

---

## Quick Start (Docker)

Run the entire stack with Docker Compose (Compose v2):

```bash
# build and run everything
docker compose up --build

# or run detached
docker compose up -d --build

# stop and remove containers
docker compose down

# stop and remove containers and volumes (WARNING: deletes Postgres data)
docker compose down -v
```

The API will be available at `http://localhost:8080` by default.

---

## Prerequisites

- Docker Desktop (with Compose v2)
- Go 1.25+ (for running the backend locally)
- Node 18+ / npm (for running the frontend locally)

Verify:

```bash
docker --version
docker compose version
go version
node --version
npm --version
```

> Note: This project prefers the `docker compose` (no hyphen) command provided by Compose v2. If you only have the legacy `docker-compose` binary, either upgrade or use the equivalent commands.

---

## Ports & Services

- API: container `autoparts-api` -> host `:8080`
- PostgreSQL: container `autoparts-postgres` -> host `:5432`
- Redis: container `autoparts-redis` -> host `:6379`
- OpenFGA (container): exposed on host `:8082` (dev only)

Inside Docker, use service hostnames (e.g. `postgres`, `redis`, `openfga`) rather than `localhost`.

---

## Docker Compose Architecture

Below is a high-level architecture diagram for the Docker Compose setup showing how services communicate inside the `autoparts-network` and how the host/browser interacts with them.

```mermaid
flowchart LR
    Browser["Browser / Host"]
    Frontend["autoparts-frontend\n(Angular)\n:4200 or :80"]
    API["autoparts-api\n(Go Gin)\n:8080"]
    Postgres["autoparts-postgres\nPostgreSQL\n:5432"]
    Redis["autoparts-redis\nRedis\n:6379"]
    OpenFGA["openfga\nOpenFGA\n:8080 (container)\n:8082 (host)"]
    Volumes["Volumes\n(postgres data, migrations)"]

    subgraph "Docker Network (autoparts-network)"
        Frontend
        API
        Postgres
        Redis
        OpenFGA
    end

    Browser -->|"HTTP (dev) :4200 / prod :80"| Frontend
    Browser -->|"HTTP :8080"| API
    Frontend -->|"HTTP API calls /api/v1"| API
    API -->|"SQL (psql)"| Postgres
    API -->|"Cache / session"| Redis
    API -->|"Authorization (OpenFGA HTTP API)"| OpenFGA
    Postgres -->|"Persisted data"| Volumes

    classDef infra fill:#f3f4f6,stroke:#333,stroke-width:1px;
    classDef db fill:#fffbeb,stroke:#b45309,stroke-width:1px;
    classDef cache fill:#ecfeff,stroke:#0369a1,stroke-width:1px;

    class Frontend,API,OpenFGA infra;
    class Postgres db;
    class Redis cache;
```

Notes:

- The host can access services using published ports (API :8080, OpenFGA :8082 development, Postgres :5432, Redis :6379).
- Containers should reference other services via Docker Compose service names (e.g. `postgres`, `redis`, `openfga`).
- Persistent data (Postgres) is stored in Docker volumes mounted by the `autoparts-postgres` service.

---

## API Endpoints

Common API endpoints exposed by the backend (prefix: `/api/v1`):

- `GET /api/v1/health` — simple health check returning `{ "status": "ok" }`.
- `POST /api/v1/auth/login` — authenticate user, returns JWT and refresh token.
- `POST /api/v1/auth/refresh` — exchange refresh token for new access token.
- `GET /api/v1/vehicles` — list vehicles for the authenticated user.
- `GET /api/v1/vehicles/{id}` — fetch a single vehicle by ID.
- `GET /api/v1/parts/search?vehicle_id={id}&query={q}` — search parts for a vehicle (used by the frontend).
- `GET /api/v1/customers` — list customers (subject to auth/authorization).
- `GET /api/v1/users` — list users (admin only in many setups).

Notes:

- Most endpoints require an Authorization header: `Authorization: Bearer <access_token>`.
- Exact request/response shapes are defined in the `backend/internal` DTOs and handler code; consult the handlers in `backend/internal/handler` for details.
- Authorization is enforced via OpenFGA policies — ensure the OpenFGA store and model IDs are configured in env/config.


## Local Development

Run the backend locally (outside Docker) for iterative development:

```bash
cd backend
# ensure your environment variables are set (see config section)
go run ./cmd/api
```

Useful Make / Taskfile targets (see `backend/Makefile` or `Taskfile.yml`):

```bash
# run tests
cd backend && go test ./...

# build binary
cd backend && go build ./cmd/api
```

## Frontend (local)

Run the Angular frontend locally for iterative UI work:

```bash
cd frontend
npm ci
# start the dev server (Angular CLI)
npm start
# or with the Angular CLI directly
npx ng serve --open

# build production bundle
npm run build
```

If the frontend needs to call the local API, update the API base URL in the frontend environment files (for example `frontend/src/environments/environment.ts`) to point to `http://localhost:8080` (or the API host you are using). The build output will be placed in `dist/` (e.g. `dist/<project-name>`).

Frontend unit tests:

```bash
cd frontend
npm test
```

Migrations are located in `backend/migrations`. To run migrations locally, use the project migration helper or psql when Postgres is available.

---

## Configuration & Environment

Configuration defaults are in `backend/configs/development.yaml` and loaded via Viper. Important environment variables (overrides):

- `DATABASE_HOST` (default: `postgres` when running with Docker)
- `DATABASE_USER`, `DATABASE_PASSWORD`, `DATABASE_NAME`
- `REDIS_HOST` (default: `redis`)
- `OPENFGA_APIURL` (container: `http://openfga:8080`), `OPENFGA_STOREID`, `OPENFGA_AUTHORIZATIONMODELID`
- `JWT_SECRET` (production only)

Example minimal env file for local Docker development (create `.env` or `docker-compose.override.yml`):

```env
DATABASE_HOST=postgres
DATABASE_USER=postgres
DATABASE_PASSWORD=postgres
DATABASE_NAME=autoparts
REDIS_HOST=redis
OPENFGA_APIURL=http://openfga:8080
OPENFGA_STOREID=01KZW369WB6CPEY6RBR42NWF1B
OPENFGA_AUTHORIZATIONMODELID=01KZW3ZKH129S283DVN5JD501G
```

---

## OpenFGA

The project integrates OpenFGA for fine-grained authorization. When running via Docker Compose the API container talks to OpenFGA at `http://openfga:8080`. From the host machine OpenFGA is reachable at `http://localhost:8082` (development). Do not use host `localhost` inside containers.

OpenFGA config (development) lives in `backend/configs/development.yaml`. The OpenFGA client code is in `backend/internal/authz/openfga.go`.

---

```mermaid
flowchart LR
    subgraph "Docker Network"
        API["autoparts-api\n:8080"]
        OpenFGA["openfga\n:8080"]
        Postgres["autoparts-postgres\n:5432"]
        Redis["autoparts-redis\n:6379"]
    end

    Browser["Host / Developer"] -->|"HTTP (dev) :8082"| OpenFGA
    Browser -->|"HTTP :8080"| API
    API -->|"gRPC/HTTP -> OpenFGA API"| OpenFGA
    API -->|"SQL"| Postgres
    API -->|"Cache"| Redis

    classDef services fill:#f9f,stroke:#333,stroke-width:1px;
    class API,OpenFGA,Postgres,Redis services;
```

### Quick OpenFGA verification (curl)

Set these environment variables (replace values as needed):

```bash
export OPENFGA_APIURL=http://localhost:8082
export OPENFGA_STOREID=01KZW369WB6CPEY6RBR42NWF1B
export OPENFGA_AUTHORIZATIONMODELID=01KZW3ZKH129S283DVN5JD501G
```

- List stores (should include your `OPENFGA_STOREID`):

```bash
curl -s "$OPENFGA_APIURL/stores" | jq .
```

- Run a permission check against the store (example):

```bash
curl -s -X POST "$OPENFGA_APIURL/stores/$OPENFGA_STOREID/check" \
    -H "Content-Type: application/json" \
    -d "{\"authorization_model_id\":\"$OPENFGA_AUTHORIZATIONMODELID\",\"tuple_key\":{\"object\":\"vehicle:vehicle-1\",\"relation\":\"viewer\",\"user\":\"user:alice\"}}" | jq .
```

If the `allowed` field in the response is `true` the check passed for the given tuple; if `false`, the tuple is not present/allowed.

### OpenFGA troubleshooting checklist

1. Confirm OpenFGA is reachable from the API container:

```bash
docker exec -it autoparts-api curl -s --fail "$OPENFGA_APIURL/health" || echo "OpenFGA not reachable from API container"
```

2. Ensure the `OPENFGA_STOREID` appears in the stores list:

```bash
curl -s "$OPENFGA_APIURL/stores" | jq -r '.[].id' | grep "$OPENFGA_STOREID" || echo "Store not found"
```

3. Verify an authorization model id exists and matches `OPENFGA_AUTHORIZATIONMODELID` (inspect store models or use the OpenFGA tooling).

4. Run the permission check example above and confirm `allowed` is `true` for a known tuple.

5. Check API logs for OpenFGA-related errors:

```bash
docker compose logs -f api | grep -i openfga
```

### Validation script

There is a small convenience script at `backend/scripts/validate_openfga.sh` that performs the checks above. It requires `curl` and optionally `jq` for nicer output.

---

## Running Tests

Backend unit tests (Go):

```bash
cd backend
go test ./...
```

Frontend tests (Angular):

```bash
cd frontend
npm ci
npm test
```

Docker smoke test (quick stack verify):

```bash
./backend/scripts/docker-smoke-test.sh
```

---

## Connecting to PostgreSQL

You can connect to the Postgres instance in several ways depending on whether you are on the host or inside Docker.

- From the host (when ports are published):

```bash
# interactive psql using host/port
psql -h localhost -p 5432 -U postgres -d autoparts

# or using a connection URL
psql "postgresql://postgres:postgres@localhost:5432/autoparts?sslmode=disable"
```

- From the Postgres container:

```bash
docker exec -it autoparts-postgres psql -U postgres -d autoparts
```

- From another container on the same Docker network (e.g. the API container), use the service hostname `postgres`:

```bash
# example: run psql inside the api container (if psql is available)
docker exec -it autoparts-api sh -c 'psql postgresql://postgres:postgres@postgres:5432/autoparts'
```

Common `psql` commands:

- `\dt` — list tables
- `\d <table>` — describe table schema
- `\q` — quit

Connection string examples (env vars):

```
DATABASE_HOST=postgres
DATABASE_USER=postgres
DATABASE_PASSWORD=postgres
DATABASE_NAME=autoparts
DATABASE_URL=postgres://postgres:postgres@postgres:5432/autoparts?sslmode=disable
```

## Connecting to Redis

- From the Redis container:

```bash
docker exec -it autoparts-redis redis-cli
```

- From the host (when ports are published):

```bash
redis-cli -h localhost -p 6379
```

- Basic Redis commands to verify connectivity:

```
PING
# expected: PONG
SET foo bar
GET foo
```

- If Redis is configured with a password, provide it with `-a` or in the URL:

```bash
redis-cli -h localhost -p 6379 -a yourpassword
# or redis://:yourpassword@localhost:6379/0
```

- From an application container, use the hostname `redis` and port `6379` (env vars):

```
REDIS_HOST=redis
REDIS_PORT=6379
# REDIS_PASSWORD=...
```

Warning: avoid `FLUSHALL` on production databases — it deletes all data.

---

## Troubleshooting

- Database connection errors: ensure `autoparts-postgres` is running and that `DATABASE_HOST` matches the container hostname `postgres`.
- Migrations: if migrations don't run on startup, run them manually using the migration helper or `psql` inside the Postgres container.
- OpenFGA: if authorization checks fail, confirm `OPENFGA_APIURL` and `STOREID`/`AUTHORIZATIONMODELID` values in environment/config.
- Port conflicts: ensure no local process occupies ports 8080/5432/6379.

Helpful commands:

```bash
# view container logs
docker compose logs -f api

# run a Postgres psql shell
docker exec -it autoparts-postgres psql -U postgres -d autoparts
```

---

## Contributing

- Follow the preferred branch/PR workflow used by the project (feature branches, small PRs).
- Run `go test ./...` and `npm test` before opening a PR.
- Add or update migrations in `backend/migrations` to match schema changes.

If you want, I can add a CONTRIBUTORS.md with a suggested workflow and checklist.

---

## License

See the `LICENSE` file in the project root.
