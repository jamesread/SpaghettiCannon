FROM alpine

LABEL org.opencontainers.image.source https://github.com/jamesread/SpaghettiCannon

ADD config.yaml /config
ADD SpaghettiCannon /SpaghettiCannon
ADD webui /webui

VOLUME /data
VOLUME /config

ENTRYPOINT ["/SpaghettiCannon"] 
