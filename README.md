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

たとえばWebサーバを更新する必要があったとする

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

## license

たぶん MIT になる。

---

This software includes the work that is distributed in the Apache License 2.0.