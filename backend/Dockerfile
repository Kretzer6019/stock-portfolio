FROM golang:1.19-alpine

WORKDIR /app/backend
COPY go.mod ./
COPY go.sum ./
RUN go mod download
COPY *.go ./

RUN go build backend