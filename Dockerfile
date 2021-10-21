FROM golang:1.15-alpine as golang

RUN apk add --no-cache bash

# Move to working directory /app
WORKDIR /app

COPY . ./

RUN go mod download
RUN go build -o /multiplica cmd/main.go

EXPOSE 8080
EXPOSE 9090

ENTRYPOINT ["/multiplica"]
