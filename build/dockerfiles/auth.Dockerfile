FROM golang as builder

WORKDIR /app
COPY ./ /app

#RUN go get github.com/githubnemo/CompileDaemon
#RUN go build cmd/main_server.go
#CMD ["go run","cmd/main_server.go"]
#ENTRYPOINT CompileDaemon --build="go build cmd/main_server.go" --command=./cmd/cmd
RUN CGO_ENABLED=0 go build -o auth_service cmd/auth/auth_service.go

FROM alpine
WORKDIR /app
COPY --from=builder /app/auth_service /app/
COPY --from=builder /app/config/configuration.json /app/
RUN chmod +x /app/auth_service
ENTRYPOINT /app/auth_service