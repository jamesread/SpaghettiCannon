FROM golang:1.26-alpine AS sqlmigrate
RUN go install github.com/rubenv/sql-migrate/sql-migrate@v1.8.1

FROM alpine

LABEL org.opencontainers.image.source=https://github.com/jamesread/SpaghettiCannon

COPY --from=sqlmigrate /go/bin/sql-migrate /usr/bin/sql-migrate
COPY config.yaml /config/config.yaml
COPY SpaghettiCannon /SpaghettiCannon
COPY frontend/dist /webui
COPY database /var/app/database
COPY docker-entrypoint.sh /usr/local/bin/spaghetticannon-entrypoint.sh
RUN chmod +x /usr/local/bin/spaghetticannon-entrypoint.sh /SpaghettiCannon /usr/bin/sql-migrate

VOLUME /data
VOLUME /config

EXPOSE 8080

ENTRYPOINT ["/usr/local/bin/spaghetticannon-entrypoint.sh"]
