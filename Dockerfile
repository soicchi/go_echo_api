FROM golang:1.21.0-bullseye AS dev
WORKDIR /app
COPY ./src/go.mod ./src/go.sum ./
RUN go mod download
COPY ./src .
CMD ["go", "run", "./src/cmd/go_echo_api/main.go"]
