FROM golang:latest

WORKDIR /go/src/github.com/hobbyfarm/gargantua
COPY . .

ENV GOOS=linux 
ENV CGO_ENABLED=0

RUN go get -d -v ./...
