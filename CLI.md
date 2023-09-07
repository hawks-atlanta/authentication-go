# CLI

This document describes how to use the service as a CLI tool.

## Environment variables

| Variable          | Description                                                  | Example                                                      |
| ----------------- | ------------------------------------------------------------ | ------------------------------------------------------------ |
| `JWT_SECRET`      | Secret used by the JWT engine                                | `CAPY_FILE`                                                  |
| `DATABASE_ENGINE` | Database engine to use. Available are `postgres` and `sqlite` | `postgres`                                                   |
| `DATABASE_DSN`    | Database connection DSN to use                               | `host=127.0.0.1 user=sulcud password=sulcud dbname=sulcud port=5432 sslmode=disable` |

## Flags

```shell
authentication :8080
```
