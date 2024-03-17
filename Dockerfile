FROM alpine

ADD config.yaml /config
ADD SpaghettiCannon /SpaghettiCannon
ADD webui /webui

VOLUME /data
VOLUME /config

ENTRYPOINT ["/SpaghettiCannon"] 
