FROM golang:1.20-alpine AS build

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY ./ ./

RUN CGO_ENABLED=0 GOOS=linux go build -o /web3

FROM alpine:latest
WORKDIR /app

RUN apk update && apk add --no-cache bash

COPY --from=build /web3 /web3

EXPOSE 8080

CMD ["/web3"]