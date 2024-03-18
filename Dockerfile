FROM alpine

LABEL org.opencontainers.image.source https://github.com/jamesread/SpaghettiCannon

ADD config.yaml /config/config.yaml
ADD SpaghettiCannon /SpaghettiCannon
ADD webui /webui

VOLUME /data
VOLUME /config

ENTRYPOINT ["/SpaghettiCannon"]
