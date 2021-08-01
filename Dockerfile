FROM golang:1.16-alpine

WORKDIR /app

COPY . ./

RUN go mod download

ENV CGO_ENABLED 0 

CMD go test ./... -v