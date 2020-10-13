FROM golang

WORKDIR /app
COPY ./ /app

ENTRYPOINT go run cmd/main_server.go
