# Autoparts-Pro
# Running AutoParts Pro with Docker

## Prerequisites

Before running AutoParts Pro with Docker, ensure you have:

* Docker Desktop installed
* Docker Compose v2 installed
* Go 1.25+ installed (for local development)

Verify Docker:

```bash
docker --version

docker-compose version
```

---

# Docker Architecture

The Docker environment runs:

| Service           | Container            | Port   | Purpose                 |
| ----------------- | -------------------- | ------ | ----------------------- |
| AutoParts Pro API | `autoparts-api`      | `8080` | Go Gin REST API         |
| PostgreSQL        | `autoparts-postgres` | `5432` | Application database    |
| Redis             | `autoparts-redis`    | `6379` | Cache / session storage |

The services communicate using the Docker network:

```
autoparts-network
```

---

# Start the Application

From the project root:

```bash
docker-compose up --build
```

The first build will:

1. Download Go dependencies
2. Build the Go API binary
3. Copy application configuration
4. Copy database migrations
5. Start PostgreSQL
6. Start Redis
7. Run database migrations
8. Start the API server

Expected API startup:

```
Loaded configuration: development
AutoParts Pro API started
Listening and serving HTTP on :8080
```

---

# Run in Background

To start containers detached:

```bash
docker-compose up -d --build
```

View logs:

```bash
docker-compose logs -f api
```

---

# Verify Container Health

Check running containers:

```bash
docker ps
```

Expected:

```
NAME                  STATUS

autoparts-postgres    Up (healthy)

autoparts-redis       Up (healthy)

autoparts-api         Up (healthy)
```

---

# API Health Check

Verify the API:

```bash
curl http://localhost:8080/api/v1/health
```

Expected response:

```json
{
  "status": "ok"
}
```

---

# Docker Smoke Test

AutoParts Pro includes a Docker smoke test.

Run:

```bash
./backend/scripts/docker-smoke-test.sh
```

Expected:

```
Checking API health...

API healthy

{"status":"ok"}

Docker smoke test passed
```

---

# Stop the Application

Stop containers:

```bash
docker-compose down
```

Remove containers and volumes:

```bash
docker-compose down -v
```

**Warning:** Removing volumes deletes the PostgreSQL database data.

---

# Rebuild After Code Changes

Rebuild the API image:

```bash
docker-compose build api
```

Restart:

```bash
docker-compose up
```

For a clean rebuild:

```bash
docker-compose build --no-cache
docker-compose up
```

---

# Environment Configuration

The API loads configuration from:

```
backend/configs/development.yaml
```

Docker overrides configuration values using environment variables.

Example:

```yaml
database:
  host: postgres
```

Docker override:

```yaml
DATABASE_HOST=postgres
```

Important Docker networking rules:

| Component  | Docker Hostname |
| ---------- | --------------- |
| PostgreSQL | `postgres`      |
| Redis      | `redis`         |
| API        | `api`           |

Do not use `localhost` between containers.

---

# Database Access

Connect to PostgreSQL:

```bash
docker exec -it autoparts-postgres psql \
-U postgres \
-d autoparts
```

List tables:

```sql
\dt
```

Exit:

```sql
\q
```

---

# Redis Access

Connect:

```bash
docker exec -it autoparts-redis redis-cli
```

Test:

```bash
ping
```

Expected:

```
PONG
```

---

# Running Tests

Run Go tests locally:

```bash
cd backend

go test ./...
```

Expected:

```
ok github.com/megatr0n28/autoparts-pro/backend/tests/smoke
```

---
## OpenFGA Authorization

AutoParts Pro uses [OpenFGA](https://openfga.dev/) for fine-grained authorization.

OpenFGA is responsible for determining whether a user has permission to access
resources such as vehicles, customers, orders, and invoices.

### OpenFGA Architecture

When running with Docker Compose:

```text
                         Docker Network
                              |
             +----------------+----------------+
             |                                 |
             v                                 v
       autoparts-api                       openfga
          :8080                             :8080
             |                                 ^
             |                                 |
             +---- http://openfga:8080 --------+
```
---
* The API container communicates with OpenFGA using:
`http://openfga:8080`
* From the host machine, OpenFGA is exposed through:
http://localhost:8082
* Do not use localhost:8082 from inside the API container.
* Inside Docker, localhost refers to the API container itself. The Docker
service name openfga must be used for container-to-container communication.
#### OpenFGA Store
The AutoParts Pro OpenFGA store is:
```sh
Store ID:
01KZW369WB6CPEY6RBR42NWF1B
```
The current authorization model is:
```sh
Authorization Model ID:
01KZW3ZKH129S283DVN5JD501G
```
These values are development environment configuration values.
#### OpenFGA Configuration
The OpenFGA configuration is defined in:
```sh
backend/internal/config/config.go
```
The configuration structure is:
```go
type OpenFGAConfig struct {
    APIURL               string `mapstructure:"apiurl"`
    StoreID              string `mapstructure:"storeid"`
    AuthorizationModelID string `mapstructure:"authorizationmodelid"`
}
```
The application configuration contains:
```go
type Config struct {
    App      AppConfig      `mapstructure:"app"`
    Log      LogConfig      `mapstructure:"log"`
    Database DatabaseConfig `mapstructure:"database"`
    Redis    RedisConfig    `mapstructure:"redis"`
    JWT      JWTConfig      `mapstructure:"jwt"`
    OpenFGA  OpenFGAConfig  `mapstructure:"openfga"`
}
```
#### Development YAML Configuration
The development configuration is located at:
```sh
backend/configs/development.yaml
```
The OpenFGA section should contain:
```YAML
openfga:
  apiurl: "http://openfga:8080"
  storeid: "01KZW369WB6CPEY6RBR42NWF1B"
  authorizationmodelid: "01KZW3ZKH129S283DVN5JD501G"
```
#### Docker Environment Variables
The API container receives the OpenFGA configuration through environment
variables:
```YAML
OPENFGA_APIURL: http://openfga:8080
OPENFGA_STOREID: 01KZW369WB6CPEY6RBR42NWF1B
OPENFGA_AUTHORIZATIONMODELID: 01KZW3ZKH129S283DVN5JD501G
```
Viper maps these environment variables to the OpenFGA configuration using:
```GO
v.SetEnvKeyReplacer(
    strings.NewReplacer(".", "_"),
)

v.AutomaticEnv()
```
Therefore:
```text
openfga.apiurl
        |
        +--> OPENFGA_APIURL

openfga.storeid
        |
        +--> OPENFGA_STOREID

openfga.authorizationmodelid
        |
        +--> OPENFGA_AUTHORIZATIONMODELID
```
#### OpenFGA Client
The OpenFGA client is implemented in:
```text
backend/internal/authz/openfga.go
```
The client provides:
```GO
Check(
    ctx context.Context,
    user string,
    relation string,
    object string,
)
```
for authorization checks.
It also provides:
```Go
CheckConnection(
    ctx context.Context,
)
```
to verify connectivity to OpenFGA.
The client is initialized during application startup.
#### Verify OpenFGA
Check the OpenFGA service:
```bash
curl http://localhost:8082/stores
```
A successful response should contain the AutoParts Pro store:
```JSON
{
  "stores": [
    {
      "id": "01KZW369WB6CPEY6RBR42NWF1B",
      "name": "autoparts-pro"
    }
  ],
  "continuation_token": ""
}
```
#### Verify Authorization Model
List the authorization models:
```bash
curl \
  "http://localhost:8082/stores/01KZW369WB6CPEY6RBR42NWF1B/authorization-models"
```
The response should contain:
```text
01KZW3ZKH129S283DVN5JD501G
```
as the current authorization model ID.

#### Verify Docker Services
Check the running services:
```bash
docker compose ps
```
Expected services include:
```text
autoparts-api
autoparts-frontend
autoparts-postgres
autoparts-redis
openfga
```
View API logs:
```bash
docker compose logs -f api
```
View OpenFGA logs:
```bash
docker compose logs -f openfga
```
#### Rebuild After Configuration Changes
After modifying OpenFGA configuration or Go source:
```bash
docker compose down
docker compose up -d --build
```
Check the API:
```bash
curl http://localhost:8080/api/v1/health
```
Expected response:
```json
{
  "status": "ok"
}
```
Check OpenFGA:
```bash
curl http://localhost:8082/stores
```
---
## OpenFGA Troubleshooting
```text
relation "store" does not exist
```
If OpenFGA reports:
```text
ERROR: relation "store" does not exist
```
- OpenFGA is connected to a PostgreSQL database that has not been initialized
with the OpenFGA schema.
- The application database and OpenFGA database should not be treated as the
same logical database unless OpenFGA has explicitly initialized its schema.
- Check the OpenFGA PostgreSQL configuration and initialization.
---
404 page not found on port 8080
If:
```bash
curl http://localhost:8080/stores
```
returns:
```text
404 page not found
```
you are likely calling the AutoParts Pro API rather than OpenFGA.
AutoParts Pro uses:
```text
localhost:8080
```
OpenFGA uses:
```text
localhost:8082
```
---
`Internal Server Error` from `/stores`

If:
```bash
curl http://localhost:8081/stores
```
returns:
```json
{
  "code": "internal_error",
  "message": "Internal Server Error"
}
```
verify that the request is using the correct OpenFGA HTTP port.

OpenFGA's gRPC and HTTP ports are separate.

The current host configuration uses:
```text
OpenFGA HTTP: localhost:8082
OpenFGA gRPC: localhost:8081
```
The REST API should be tested with:
```bash
curl http://localhost:8082/stores
```
---
## Authorization Roadmap
OpenFGA will eventually be used for relationships such as:
```text
user -> customer
customer -> vehicle
customer -> invoice
customer -> order
```
For example:
```text
user:123
    |
    +-- owns --> vehicle:456
```
The application can then ask OpenFGA:
```text
Can user:123 view vehicle:456?
```
The expected authorization flow is:
```text
Angular Frontend
       |
       v
AutoParts Pro API
       |
       v
OpenFGA Authorization Check
       |
       +---- allowed ----> access resource
       |
       +---- denied ------> HTTP 403
```
---
The next implementation step is to define the AutoParts Pro authorization
model and write the initial customer-to-vehicle ownership tuples.

### One important README correction

Your current ports should be documented as:
```text
| Service | Host | Container |
|---|---:|---:|
| AutoParts API | `8080` | `8080` |
| OpenFGA HTTP | `8082` | `8080` |
| OpenFGA gRPC | `8081` | `8081` |
| PostgreSQL | `5432` | `5432` |
| Redis | `6379` | `6379` |
| Angular | `4200` | `4200` |
```
The **API → OpenFGA** connection is:

```text
http://openfga:8080
```
while your Mac → OpenFGA connection is:
```text
http://localhost:8082
```
---
# Troubleshooting

## API cannot connect to PostgreSQL

Check:

```bash
docker-compose ps
```

PostgreSQL should show:

```
healthy
```

View logs:

```bash
docker-compose logs postgres
```

---

## Migration Error

Example:

```
failed to open source "file://migrations"
```

Verify the image contains migrations:

```bash
docker exec -it autoparts-api ls migrations
```

---

## Configuration Error

Example:

```
Config File "development" Not Found
```

Verify:

```bash
docker exec -it autoparts-api ls configs
```

Expected:

```
development.yaml
production.yaml
app.yaml
```

---

# Development Workflow

Recommended workflow:

1. Make code changes
2. Run tests:

```bash
go test ./...
```

3. Rebuild Docker:

```bash
docker-compose build api
```

4. Restart:

```bash
docker-compose up
```

5. Run smoke test:

```bash
./backend/scripts/docker-smoke-test.sh
```

---

# Current Docker Support

AutoParts Pro currently supports:

* Go 1.25 Docker builds
* PostgreSQL 17
* Redis 7
* Automated migrations
* API health checks
* Container health checks
* Docker smoke testing


# Troubleshooting Guide

This section contains common issues encountered during AutoParts Pro backend development and their solutions.

---

## PostgreSQL Troubleshooting

## PostgreSQL Connection Refused

#### Error
```
failed to connect to user=postgres database=
localhost:5432: connect: connection refused
```

#### Cause

PostgreSQL container is not running or API is pointing to the wrong database host.

### Check running containers

```bash
docker ps
```

Expected
```bash
postgres
5432->5432
```

Password Authentication Failed
Error
```bash
FATAL: password authentication failed for user "postgres"
```
Verify credentials
Check:
```bash
configs/development.yaml
```
Example:
```bash
database:
  host: localhost
  port: 5432
  name: autoparts
  user: postgres
  password: postgres
```
Connected to Wrong Database
Symptoms
```bash
\dt

Did not find any relations.
```
Expected:
```bash
autoparts
```
If connected to:
```bash
postgres
```
switch:
```bash
\c autoparts
```

## Database Migration Troubleshooting
Relation Already Exists
#### Error
```bash
relation "customers" already exists
```

#### Cause
- Migration was executed more than once.
Check migrations
```sql
SELECT *
FROM schema_migrations;
```

Check tables
```sql
\dt
```
If the database is disposable:
```sql
DROP DATABASE autoparts;
CREATE DATABASE autoparts;
```
Then rerun migrations.

## Docker PostgreSQL Access
Enter PostgreSQL container:
```bash
docker exec -it autoparts-postgres psql -U postgres
```
Connect to AutoParts database:
```sql
\c autoparts
```
List tables:
```sql
\dt
```

## JWT Authentication Troubleshooting
Invalid Token
#### Error
```bash
{
 "error":"invalid token"
}
```
#### Cause
- Using placeholder token:
```bash
YOUR_ACCESS_TOKEN
```
instead of a real JWT.
Login first:
```bash
curl -X POST \
http://localhost:8080/api/v1/auth/login \
-H "Content-Type: application/json" \
-d '{
"email":"admin@test.com",
"password":"Password123"
}'
```
Copy:
```bash
access_token
```
Use:
```bash
curl \
http://localhost:8080/api/v1/users/me \
-H "Authorization: Bearer ACCESS_TOKEN"
```

## Refresh Token Troubleshooting
Invalid Refresh Token
#### Error
```bash
{
 "error":"invalid refresh token"
}
```

Check stored tokens
```sql
SELECT
id,
user_id,
token_hash,
expires_at,
revoked
FROM refresh_tokens;
```
Refresh Token Immediately Revoked
Symptoms
Database shows:
```
revoked = true
```
or:
```
id = 00000000-0000-0000-0000-000000000000
```
#### Cause
UUID was not generated.
#### Fix
Refresh token model must generate UUID:
```go
func (r *RefreshToken) BeforeCreate(
tx *gorm.DB,
) error {

if r.ID == uuid.Nil {
r.ID = uuid.New()
}

return nil
}
```

Go Compilation Errors
Undefined jwtManager
#### Example:
```
undefined: jwtManager
```
#### Cause
Dependency was not initialized before router creation.
Verify:
```
bootstrap/application.go
```
Order should be:
```
Config
 |
Logger
 |
Database
 |
Repositories
 |
JWT Manager
 |
Services
 |
Handlers
 |
Router
```
### Undefined Handler
#### Example:
```
undefined: userHandler
```
#### Cause
Handler was not created before passing to router.
#### Example:
```go
userHandler :=
handler.NewUserHandler()
```

## Import Collision Problems
os/user.User Error
#### Example:
```
unknown field FirstName in struct literal of type os/user.User
```
#### Cause
Wrong package imported.
Incorrect:
```go
import "os/user"
```
Correct:
```go
import domainUser "github.com/megatr0n28/autoparts-pro/backend/internal/domain/user"
```

## Gin Troubleshooting
API Starts But curl Fails
#### Error
```
curl: Failed to connect to localhost port 8080
```
Check API process
```bash
go run ./cmd/api
```
Expected:
```
AutoParts Pro API started
```
Check port:
```bash
lsof -i :8080
```

## Customer Profile Troubleshooting
Customer Profile Not Found
#### Response
```json
{
"error":"customer profile not found"
}
```
#### Cause
User exists but profile was created before automatic profile creation was added.
Verify:
```sql
SELECT *
FROM customer_profiles;
```

## Configuration Troubleshooting
Loaded Configuration
Startup should show:
```
Loaded configuration: development
```
Configuration files:
```
configs/
├── development.yaml
└── production.yaml
```
Environment override:
```bash
export APP_ENV=production
```

## Git Troubleshooting
Check changes:
```bash
git status
```
Commit:
```bash
git add .

git commit -m "message"

git push origin main
```
View history:
```bash
git log --oneline
```

## Recommended Debug Checklist
When the API fails:
1. Verify PostgreSQL is running
```bash
docker ps
```
2. Verify database
```sql
SELECT current_database();
```
3. Verify migrations
```sql
\dt
```
4. Start API
```bash
go run ./cmd/api
```
5. Test health endpoint
```bash
curl http://localhost:8080/api/v1/health
```
6. Login
```bash
curl -X POST /api/v1/auth/login
```
7. Test protected endpoint
```bash
curl /api/v1/users/me
```
