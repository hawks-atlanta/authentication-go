FROM hashicorp/http-echo:latest

EXPOSE 8080/tcp

ENV DATABASE_DSN = ""

ENTRYPOINT ["/http-echo", "--listen", ":8080", "--text", "DATABASE_DSN=${DATABASE_DSN}"]