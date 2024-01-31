FROM golang:1.20-alpine AS build

WORKDIR /

COPY . .

RUN go mod download
RUN go build -o /hash-service cmd/refresh-hash/main.go

FROM alpine:3.14

WORKDIR /

COPY --from=build ./hash-service ./hash-service

EXPOSE 8080

CMD [ "/hash-service" ]
