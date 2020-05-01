FROM alpine:latest
COPY ./bin/promethues-webhook-wechatwork /
ENTRYPOINT ["/promethues-webhook-wechatwork"]