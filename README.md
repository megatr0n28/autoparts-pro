# autoparts-pro

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
