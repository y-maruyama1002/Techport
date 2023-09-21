FROM golang:1.19.13-alpine3.18

WORKDIR /app

COPY src/go.mod .
COPY src/go.sum .

RUN go get -u github.com/cosmtrek/air && go build -o /go/bin/air github.com/cosmtrek/air
CMD ["air", "-c", ".air.toml"]
