FROM golang:1.11.2-alpine3.8 AS build

WORKDIR /
COPY . /go/src/github.com/t0w4/toDoListBackend
RUN apk update \
  && apk add --no-cache git \
  && go get github.com/go-sql-driver/mysql \
  && go get github.com/google/uuid \
  && go get github.com/gorilla/mux
RUN cd /go/src/github.com/t0w4/toDoListBackend && go build -o bin/todolist main.go

FROM alpine:3.8
COPY --from=build /go/src/github.com/t0w4/toDoListBackend/bin/todolist /usr/local/bin/
CMD ["todolist"]