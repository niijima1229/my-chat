# 使用するGolangのイメージを指定する
FROM golang:1.19
ENV GO111MODULE=on
ENV PATH=$PATH:/usr/local/go/bin
ENV GOPATH=$HOME/go/package:$HOME/go/workspace
ENV PATH=$HOME/go/package/bin:$HOME/go/workspace/bin:$PATH

WORKDIR /api

# 必要なパッケージなどなどをインストールする
RUN apt-get update \
	&& apt-get install git

RUN go install github.com/cosmtrek/air@latest

CMD ["air", "-c", ".air.toml"]

EXPOSE 8080