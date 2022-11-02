# syntax=docker/dockerfile:1

FROM golang:1.19-alpine

WORKDIR /app

COPY go.mod ./
COPY go.sum ./

COPY / /app/

RUN go mod download

RUN go build -o /apk-runtime-api-v1

CMD [ "/apk-runtime-api-v1" ]
