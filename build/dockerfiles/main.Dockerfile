FROM golang

WORKDIR /app
COPY ./ /app

#RUN go get github.com/githubnemo/CompileDaemon
RUN go run cmd/main_server.go
#ENTRYPOINT CompileDaemon --build="go build cmd/main_server.go" --command=./cmd/cmd
