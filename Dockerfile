FROM alpine

LABEL org.opencontainers.image.source https://github.com/jamesread/SpaghettiCannon

COPY config.yaml /config/config.yaml
COPY backend/SpaghettiCannon /SpaghettiCannon
COPY webui/dist/* /webui

VOLUME /data
VOLUME /config

ENTRYPOINT ["/SpaghettiCannon"]
