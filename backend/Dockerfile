FROM golang:1.19.13-alpine3.18

WORKDIR /app

COPY src/go.mod .
COPY src/go.sum .

RUN go mod download
