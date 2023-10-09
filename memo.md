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

### アーキテクチャ設計

ドメイン名のディレクトを作成してその中に
route, handler repository usecase が入る。
interface やモデル名とマッチさせる構造体は domain ディレクトリに入れる。
ドメイン名のディレクトリの流れは
handler→usecase→repository
の順番。これらの中で domain からモデルを都度呼び出している

### CRUD の実装

最後に c.JSON を書かないときは何も返さない 200 レスポンスになる

```go
func (u *blogUsecase) GetById(id int64) (res domain.Blog, err error) {
	res, err = u.blogRepo.GetById(id)
	if err != nil {
		return
	}
	return
}
```

みたいに、レスポンス値を変数名とセットで定義すると
関数の中身ではその変数は最初から定義された変数として扱える。
また、
ただの return だったとしてもその両変数が暗黙的に return される

gin で get のクエリパラメータを取る

```go
// localhost:3000/api/v1/blogs?num=10
num := c.Query("num")
fmt.Println(num)
// 10
```

### next の開発

npx create-next-app@latest frontend
でアプリの作成

```
$ npx create-next-app@latest frontend
Need to install the following packages:
  create-next-app@13.5.2
Ok to proceed? (y) y
✔ Would you like to use TypeScript? … No / Yes
✔ Would you like to use ESLint? … No / Yes
✔ Would you like to use Tailwind CSS? … No / Yes
✔ Would you like to use `src/` directory? … No / Yes
✔ Would you like to use App Router? (recommended) … No / Yes
✔ Would you like to customize the default import alias? … No / Yes
```

global.css は全てのページで適用される
.module.css は特定のものに割り当てるもの
npm run dev
localhost:3000 で next のランディングページが出れば成功

### next.js のアーキテクチャ設計

https://github.com/alan2207/bulletproof-react
を参考にやってみる

解説はこの辺
https://note.com/ryoppei/n/n2e3e7a66e758
https://zenn.dev/motonosuke/articles/8f4ba3714f30fe

### frontend と backend の通信

これまで、"http://172.24.0.1:3000/api/v1"を使ってなんとかGETはできていたが、POSTができなくなった。
エラーはランタイムエラー。
axios を使うと host を localhost にすると通信できる。なぜ？
とはいえ、front から back への通信で CORS 制限がかかったので、
"github.com/gin-contrib/cors"を入れて、ゆるゆるで対応

docker の network に front, back を入れるのはあまり得策じゃないのかなぁ
