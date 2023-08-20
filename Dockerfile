FROM golang:1.21.0-alpine3.18 AS builder

RUN apk --no-cache add git
WORKDIR /build
COPY go.mod /src/go.mod
COPY go.sum /src/go.sum
WORKDIR /src
RUN go mod download
COPY . /src
RUN go build -o /build/ ./...

FROM alpine:3.18 AS service

EXPOSE 8080/tcp
COPY --from=builder /build/ /service-bin/
WORKDIR /db
WORKDIR /
ENV DATABASE_ENGINE "sqlite"
ENV DATABASE_DSN    "/db/database.db?cache=shared&mode=rwc"
ENTRYPOINT ["/service-bin/authentication-go", ":8080"]
