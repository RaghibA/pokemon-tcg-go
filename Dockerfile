FROM golang:1.20.2-alpine3.17

WORKDIR /usr/src/app

COPY . . 

RUN go mod tidy
