ARG GO_VERSION=1.18
FROM golang:${GO_VERSION}-alpine

WORKDIR /jwt-filter

COPY ./go.mod .
COPY ./main.go .

RUN go get ./...

ENTRYPOINT go run main.go
