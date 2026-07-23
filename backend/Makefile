APP=autoparts-api

run:
	go run ./cmd/api

build:
	go build -o bin/$(APP) ./cmd/api

test:
	go test ./...

fmt:
	go fmt ./...

tidy:
	go mod tidy

migrate-up:
	go run ./cmd/migrate up

migrate-down:
	go run ./cmd/migrate down

docker-up:
	docker-compose up -d


docker-down:
	docker-compose down


docker-reset:
	docker-compose down -v --remove-orphans

docker-smoke:
	./scripts/docker-smoke-test.sh

docker-logs:
	docker-compose logs -f