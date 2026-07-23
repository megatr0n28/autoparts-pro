FROM golang:1.25-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux \
    go build -o autoparts-api ./cmd/api


FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/autoparts-api .

COPY configs ./configs

COPY migrations ./migrations

EXPOSE 8080

CMD ["./autoparts-api"]