FROM golang:1.17.13-bullseye

WORKDIR /opt/app

COPY . .

RUN go mod download \
 && go build ./cmd/main.go
 