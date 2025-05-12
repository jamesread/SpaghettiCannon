FROM alpine

LABEL org.opencontainers.image.source=https://github.com/jamesread/SpaghettiCannon

COPY config.yaml /config/config.yaml
COPY SpaghettiCannon /SpaghettiCannon
COPY frontend/dist /webui

VOLUME /data
VOLUME /config

ENTRYPOINT ["/SpaghettiCannon"]
