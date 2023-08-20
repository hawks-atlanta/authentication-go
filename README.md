# authentication-go

[![Release](https://github.com/hawks-atlanta/authentication-go/actions/workflows/release.yaml/badge.svg)](https://github.com/hawks-atlanta/authentication-go/actions/workflows/release.yaml)
[![Tagging](https://github.com/hawks-atlanta/authentication-go/actions/workflows/tagging.yaml/badge.svg)](https://github.com/hawks-atlanta/authentication-go/actions/workflows/tagging.yaml)
[![Go Report Card](https://goreportcard.com/badge/github.com/hawks-atlanta/authentication-go)](https://goreportcard.com/report/github.com/hawks-atlanta/authentication-go)
![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/hawks-atlanta/authentication-go)

Microservice used to authorize user inside the system.

## Documentation

|       Document        |                             URL                              |
| :-------------------: | :----------------------------------------------------------: |
|   CLI documentation   |                       [CLI.md](CLI.md)                       |
|         CICD          | [CICD.md](https://github.com/hawks-atlanta/docs/blob/main/CICD.md) |
|    Database models    | [Database.md](https://github.com/hawks-atlanta/docs/blob/main/Database.md#Authentication) |
|     CONTRIBUTING      |              [CONTRIBUTING.md](CONTRIBUTING.md)              |
| OpenAPI specification |           [Specification](docs/spec.openapi.yaml)            |

## Development

You can setup the necessary services by running:

```shell
docker compose up -d
```

Then you can run tests by:

```shell
go test -v ./...
```

For locally measuring coverage:

```shell
go test -v -covermode atomic -coverprofile ./cover.txt ./...
go tool cover -html ./cover.txt
```

### Services

| Service    | HOST:PORT      | Credentials                  |
| ---------- | -------------- | ---------------------------- |
| PostgreSQL | 127.0.0.1:5432 | `username:password@database` |

## Coverage

