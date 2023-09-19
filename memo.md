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

### gorm を使った mysql とのコネクション

```yaml
db:
  image: mysql:8.1.0
  environment:
    - MYSQL_DATABASE=root
    - MYSQL_ROOT_PASSWORD=password
  volumes:
    - "data-base:/var/lib/mysql"
  command: mysqld --default-authentication-plugin=caching_sha2_password
  ports:
    - 3306:3306
```

docker-compose で mysql のコンテナを作成

main.go に gorm をインストール

```
go get -u gorm.io/gorm
go get -u gorm.io/driver/mysql
```

コネクション情報を記載
tcp(db)の db 部分は docker のコンテナ名を指定する
docker を使っていない場合は 127.0.0.1:3306 でいけるけど、コンテナ使ってる場合はコンテナ名でないとうまくいかない

```go
Product structureを使ってテーブルをマイグレートしてくれる
db.AutoMigrate(&Product{})
// Create
db.Create(&Product{Code: "D43", Price: 200})

var product Product
// 一件取り出し
db.First(&product, 1)
fmt.Println("check the value")
fmt.Println(product.Code)
//  D42
fmt.Println(product.Price)
// 100
```

### アーキテクチャ構成

https://qiita.com/ryoh07/items/8ebac006c5294b9b3f58
を元に作成してみる
ただし、使う orm は gorm
handler ではなく controller とする

流れ
main→router→di
di では repository、interactor, controller を順に生成していき、controller の struct を返す
このとき、interactor は repository を、controller を interactor を引数にとって作成されるため、
controller は interactor を実行できて、interactor は repository を実行できるようにする準備ができる
di によって返された controller を router でパスに紐づけて実行させる
controller が実行されて、controller ではパラメータを受け取って、interactor を起動する
interactor は repository を起動する
repository はデータベースにアクセスして値を取り出して interactor に返す。
interactor は値を受け取って、dto を使って、整形する。整形したものを controller に返す
controller はそれを json 形式で返す

### go の create の api

流れは同じ。
難しかったのは gorm での挙動。
entity の struct に gorm.Model を入れるということは ID, CreatedAt, UpdatedAt, DeletedAt を指定していることの暗黙。これを使ってみせるとき、作成するときを実装する
作成するときは id を指定しなくても auto increment をしてくれる gorm の挙動を利用している

いろいろしんどくなったので参考にするリポジトリを変更する

https://gist.github.com/mpppk/609d592f25cab9312654b39f1b357c60
https://github.com/bxcodec/go-clean-arch

middleware や test に関しての実装もある。獲得されているスターの数も多い

backend2 を作成してそっちでしばらく開発を続けるようにする
