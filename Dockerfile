FROM golang:1.21.0-alpine3.18 AS builder

RUN apk --no-cache add git upx
WORKDIR /build
COPY go.mod /src/go.mod
COPY go.sum /src/go.sum
WORKDIR /src
RUN go mod download
COPY . /src
RUN go build -o /build/ ./...
RUN upx /build/*

FROM alpine:3.18 AS service

EXPOSE 8080/tcp
RUN adduser -D -u 5000 -h /opt/application -s /sbin/nologin application
USER application
COPY --from=builder /build/ /opt/application
WORKDIR /opt/application/db
WORKDIR /
ENV DATABASE_ENGINE "sqlite"
ENV DATABASE_DSN    "/opt/application/db/database.db?cache=shared&mode=rwc"
ENTRYPOINT ["/opt/application/authentication-go", ":8080"]
