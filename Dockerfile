FROM golang:1.21.0 AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o main ./cmd/go_echo_api/main.go

FROM golang:1.21.0-bullseye AS dev
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
CMD ["go", "run", "./cmd/go_echo_api/main.go"]

FROM golang:1.21.0-bullseye AS prod
WORKDIR /app
COPY --from=builder /app/main .
CMD ["./main"]