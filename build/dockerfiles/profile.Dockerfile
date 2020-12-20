FROM golang as builder

WORKDIR /app
COPY ./ /app

#RUN go get github.com/githubnemo/CompileDaemon
#RUN go build cmd/main_server.go
#CMD ["go run","cmd/main_server.go"]
#ENTRYPOINT CompileDaemon --build="go build cmd/main_server.go" --command=./cmd/cmd
RUN CGO_ENABLED=0 go build -o profile_service cmd/profile/profile_service.go

FROM alpine
WORKDIR /app
COPY --from=builder /app/profile_service /app/
RUN chmod +x /app/profile_service
ENTRYPOINT /app/profile_service