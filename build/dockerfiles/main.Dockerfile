FROM golang:latest

EXPOSE 3000

CMD ["go run","../../cmd/main_server.go"]
