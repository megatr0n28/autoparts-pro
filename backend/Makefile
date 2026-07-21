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