# CLI

This document describes how to use the service as a CLI tool.

## Environment variables

| Variable       | Description                    | Example                                                      |
| -------------- | ------------------------------ | ------------------------------------------------------------ |
| `DATABASE_DSN` | Database connection DSN to use | `host=127.0.0.1 user=sulcud password=sulcud dbname=sulcud port=5432 sslmode=disable` |

## Flags

Use `--help` for more information.

```shell
authentication --help
```

Listen address

```shell
authentication --listen :8080
```
