FROM golang

WORKDIR /app
COPY ./ /app

#RUN go get github.com/githubnemo/CompileDaemon
RUN go build cmd/main_server.go
CMD ["./cmd/cmd"]
#ENTRYPOINT CompileDaemon --build="go build cmd/main_server.go" --command=./cmd/cmd
