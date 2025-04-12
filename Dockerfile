FROM golang:1.24 AS builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o resource-service ./cmd/main

FROM alpine:3.18

WORKDIR /app
COPY --from=builder /app/resource-service .

EXPOSE 8081

ENV PG_USER=string  \
    PG_PASS =string

CMD ["./resource-service"]