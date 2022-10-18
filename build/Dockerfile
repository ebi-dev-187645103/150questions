# # goバージョン
# # https://hub.docker.com/_/golang
# FROM golang:1.17.8-alpine
# # アップデートとgitのインストール！！
# RUN apk add --update &&  apk add git
# # appディレクトリの作成
# RUN mkdir /go/src/app
# # ワーキングディレクトリの設定
# WORKDIR /go/src/app
# # ホストのファイルをコンテナの作業ディレクトリに移行
# ADD . /go/src/app

FROM golang:1.18-bullseye
# ワーキングディレクトリの設定
WORKDIR /src
# ホストのファイルをコンテナの作業ディレクトリに移行
ADD . /src

RUN go mod tidy
RUN go get -u github.com/cosmtrek/air && \
    go build -o /go/bin/air github.com/cosmtrek/air
RUN go get -u  github.com/go-delve/delve/cmd/dlv@latest && \
    go build -o /go/bin/dlv github.com/go-delve/delve/cmd/dlv
