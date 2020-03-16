FROM alpine:latest
COPY ./bin/promethues-webhook-wechatwork /usr/bin
ENTRYPOINT ["/promethuess-webhook-wechatwork"]