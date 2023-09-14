## go + gin + docker-compose で開発環境構築まで

### hello world を出す

main.go を作成する
main.go に main の func を作成して fmt.Println("hello world")を実装。
go run main.go
を実行してコンソールに表示されることを確認

### gin のインストール

公式のスタート
https://pkg.go.dev/github.com/gin-gonic/gin#section-readme

import を入れる
gin の記法に沿って func main を実装
go get -u github.com/gin-gonic/gin
を実行 →go.mod が作成される
go mod init github.com/y-maruyama1002/Techport
を実行 →go.sum が作成される（package-lock の役割らしい）
go run main.go
を実行して
localhost:8080 で開くことを確認する

### docker-compose の環境にする

```dockerfile
# go 1.19.3のalpine(軽量のlinuxのイメージ)のイメージを取り込み
FROM golang:1.19.13-alpine3.18

# ワークディレクトリにappフォルダを作成
WORKDIR /app

# そこにsrc/配下にあるgo.modとgo.sumを入れる
COPY src/go.mod .
COPY src/go.sum .

# modをダウンロードしておく
RUN go mod download
```

```yml
version: "3.8.1"

services:
  app:
    build: . # dockerfileの位置
    container_name: app
    ports:
      - "8080:8000" # 8080から8000番に繋げる
    volumes:
      - ./src:/app # src配下のものをappに入れいる？
    tty: true # コンテナが起動しっぱなしになるようにする
```

docker-compose up --build -d
を実行して
localhost:8080 で json がブラウザに表示されることを確認する

### docker + go でホットリロード

docker を都度 build しなおさないとコードの変更が反映されないのはきついので、
ホットリロードして、コードの変更がすぐに反映されるようにする
air.toml を使う。

Dockerfile 内で air をインストール。
air の初期化（docker-compose run --rm app air init）→ これで.air.toml ファイルが作成される
docker-compose up で起動して、コードを変更してみて、リロードすることで反映されるのを確認
