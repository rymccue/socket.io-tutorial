FROM golang:1.11-alpine AS build

RUN apk add --no-cache git ca-certificates

WORKDIR /www

RUN go get github.com/gorilla/mux

COPY . .

RUN go build -o app .

FROM alpine

WORKDIR /www/

COPY --from=build /www/app .

ENTRYPOINT ["./app"]
