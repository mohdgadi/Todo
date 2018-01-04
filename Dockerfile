FROM golang:latest

WORKDIR /go/src/todo
COPY * /go/src/todo/

RUN go get github.com/gorilla/mux
RUN go get github.com/mattn/go-sqlite3

RUN go build
RUN go install

EXPOSE 8000
ENTRYPOINT todo
