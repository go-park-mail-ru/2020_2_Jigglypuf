FROM golang as builder

WORKDIR /app
COPY ./ /app

#RUN go get github.com/githubnemo/CompileDaemon
#RUN go build cmd/main_server.go
#CMD ["go run","cmd/main_server.go"]
#ENTRYPOINT CompileDaemon --build="go build cmd/main_server.go" --command=./cmd/cmd
RUN go install cmd/main_server.go

FROM alpine
WORKDIR /app
COPY --from=builder /go/bin/main_server /app/
RUN chmod +x ./main_server
ENTRYPOINT ./main_server