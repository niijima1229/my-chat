# 使用するGolangのイメージを指定する
FROM golang:1.17-alpine
ENV GO111MODULE=on
ENV PATH=$PATH:/usr/local/go/bin
ENV GOPATH=$HOME/go/package:$HOME/go/workspace
ENV PATH=$HOME/go/package/bin:$HOME/go/workspace/bin:$PATH

WORKDIR /api

# 必要なパッケージなどなどをインストールする
RUN apk update \
	&& apk add --no-cache git

RUN go install github.com/cosmtrek/air@latest

RUN go get github.com/gin-gonic/gin

CMD ["air", "-c", ".air.toml"]

EXPOSE 8080