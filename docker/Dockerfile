FROM golang:1.18.2-alpine

RUN apk add --update make

COPY . /opt/src

WORKDIR /opt/src

RUN go build
