FROM golang:alpine AS build-stage

WORKDIR /app

COPY . .
RUN go mod download

RUN go build -o app ./cmd/todo-list/main.go

FROM build-stage AS run-test-stage
RUN go test -v ./...

FROM scratch AS build-release-stage
COPY --from=build-stage app config ./

ENTRYPOINT ["/app", "release"]
