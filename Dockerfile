FROM golang:1.17.1-alpine3.13

WORKDIR /go/src
COPY . .
# RUN go mod tidy

CMD ["go","run","server.go"]
