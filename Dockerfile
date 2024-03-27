FROM golang:alpine

RUN mkdir -p /todo-list
WORKDIR /todo-list
COPY . .

RUN go build -o app cmd/todo-list/main.go

ENTRYPOINT ["./app release"]
