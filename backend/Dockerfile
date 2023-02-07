FROM golang:1.19-alpine

WORKDIR /app/backend

RUN go install github.com/cosmtrek/air@latest

COPY . ./
RUN go mod download
RUN air init

CMD [ "air" ]