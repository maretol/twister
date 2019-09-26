# Twister

`[Twi]n page te[ster]` -> Twister

2つのWebページを比較する。

APIサーバとかWebサーバで破壊的更新をしなければならないときとかに確認する

## Using

実行バイナリまだ作ってないから自分でビルドするか `go run main.go` で補完してほしい

---

まず設定ファイルの土台を作ってあげる
```
$ twister config create
```
たぶん実行ディレクトリに `twisterconf.json` ってのができたはず

### example:
実際に動かしてみる

たとえばWebサーバに破壊的変更が入って更新する必要があったとする

 - URL1 : `https://old.server/`
 - URL2 : `https://new.server/`

こんな感じのアドレスの場合、twisterconf.json に
```
  "urls": [
    "http://old.server",
    "http://new.server"
  ],
  "paths": [
    "/api/test/"
  ]
```
と書いてあげる。 paths はアドレスの共通部分を書いてあげる

あとは実行したらそれぞれのアドレス、 `http://old.server/api/test` と `http://new.server/api/test` にアクセスし、結果を比較してくれる。

もし結果が同じなら"Not exist the result to show"って答えてくれる。違ったら差分を出せるようにしたい（現状は一応差分を出してくれるが完璧な状態ではない）

#### Option

表示時にURLだけだとわからないケースでもURLにタグ付けしてあげると表示時にそっちを使ってくれる

その場合、設定ファイルには

```
  "urls": [
    {
      url: "http://old.server",
      tag: "old"
    },
    {
      url: "http://new.server",
      tag: "new"
    }
  ],
```

と書いてあげればいい

### example2

configファイルのurlは非共通部分、pathsはパスの共通部分で分けているので
```
  "urls": [
    "http://api.server/v1/",
    "http://api.server/v2/"
  ],
  "paths": [
    "/api/test?q=hogehoge"
  ]
```
みたいな形で少しだけ応用してテストすることもできる。

## その他

ヘッダー、ステータスコード、ボディの3要素を比較しているが、設定ファイルのそれぞれをfalseにすると比較しなくなる。
また現状ヘッダーは比較できていない。

実行時にパラメータを設定するとconfigを上書きして比較したり、--config で任意のコンフィグファイルを渡せるようになってるはず。

## to do

 - HTTPアクセスの並列処理
 - 3つ以上のページの比較
 - header の比較
 - responseがJSONやXMLだったときはテキストとしての比較ではなくデータ各要素の比較にしたい
 - Color 対応
 - リファクタリング
 - test追加＆CI対応

## license

たぶん MIT になる。(I may apply MIT License)

---

Used Library

 - viper / https://github.com/spf13/viper
 - cobra / https://github.com/spf13/cobra
 - go-diff / https://github.com/sergi/go-diff

---

This software includes the work that is distributed in the Apache License 2.0.
