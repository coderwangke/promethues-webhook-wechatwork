# build stage
FROM golang:1.12.6-alpine3.9 AS build-env
MAINTAINER wangkehenan <wangkehenan@gmail.com>
RUN apk add --no-cache git
COPY . /src
RUN cd /src && go build -o promethues-webhook-wechatwork *.go

FROM alpine:latest
COPY --from=build-env ./src/promethues-webhook-wechatwork /
ENTRYPOINT ["/promethues-webhook-wechatwork"]
