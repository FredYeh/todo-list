FROM golang:alpine AS build-stage
WORKDIR /src
COPY . .
RUN go build -o /app ./cmd/todo-list/main.go

# FROM build-stage AS run-test-stage
# RUN go test -v ./...

FROM scratch AS build-release-stage
WORKDIR /
COPY ./config ./config
COPY --from=build-stage /app /app 
CMD ["./app", "release"]
