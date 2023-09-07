FROM golang:1.21.0-bullseye AS dev
WORKDIR /app
COPY ./src/go.mod ./src/go.sum ./
RUN go mod download
COPY ./src .
EXPOSE 8000
CMD ["go", "run", "./cmd/dev/main.go"]

FROM golang:1.21.0 AS build
WORKDIR /app
COPY ./src/go.mod ./src/go.sum ./
RUN go mod download
COPY ./src .
RUN GOOS=linux GOARCH=amd64 go build -tags lambda.norpc -o main ./cmd/go_echo_api/main.go

FROM public.ecr.aws/lambda/provided:al2 AS prd
WORKDIR /app
COPY --from=build /app/main .
ENTRYPOINT ["./main"]
