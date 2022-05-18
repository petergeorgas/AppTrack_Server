# syntax=docker/dockerfile:1

FROM golang:1.18.2-bullseye
RUN mkdir /app
ADD . /app
WORKDIR /app

RUN go mod download

EXPOSE 8008

RUN go build -o main .

CMD ["/app/main"]