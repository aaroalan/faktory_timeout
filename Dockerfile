FROM golang:1.15.7-alpine

ENV CGO_ENABLED 0
ENV GOOS 'linux'

# Update dependency repositories and add build dependencies.
RUN apk update && apk add ca-certificates curl git && rm -rf /var/cache/apk/*

RUN mkdir /worker
WORKDIR /worker
COPY . /worker
RUN go build -o $GOPATH/bin/worker /worker/app.go
