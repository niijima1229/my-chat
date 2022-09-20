# 使用するGolangのイメージを指定する
FROM golang:1.19.1-alpine
ENV GO111MODULE=on

# 必要なパッケージなどなどをインストールする
RUN apk update \
	&& apk add --no-cache git

EXPOSE 8080