FROM alpine:latest
RUN apk --no-cache add ca-certificates

WORKDIR /app
COPY build/leinao-docs-sync-bot .
COPY configs/config.toml configs/

CMD ["./leinao-docs-sync-bot"]