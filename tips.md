## サードパーティパッケージをインストール

-u オプションをつけることによって、引数で指定したパッケージの最新版をインストールすることができます。

# go.mod について

go.mod ファイルにはモジュールのルートディレクトリを示す以外にも、パッケージの依存関係を記録するという役割が存在する

# go.sum

謎の文字列 h1:xxxxxx~はチェックサムと呼ばれるもので、ビルド再現性の担保のために使われる情報
このコードを別の場所でも同様に動かすためには、当然依存パッケージである gorilla/mux も同じものをダウンロードしてくる必要があります。「確実に同じ依存パッケージを使っているのか」ということを担保するためにチェックサムの情報が必要で、これにより「どこで動かしても、同じ動きをさせることができる」ビルド再現性が実現できます。

# io.ReadCloser

http.Request.Body の中からリクエストボディの内容を読み出す処理としては、

1. 何らかの方法でバイトスライスを用意
2. Body の Read メソッドを呼び出して、1 で用意したバイトスライスに内容を書き込む
3. 使い終わった Body を Close メソッドで閉じる

## encoding/json でメモリを扱う

### json.Marshal・json.Unmarshal 関数

引数・戻り値に []byte を含んでいることから、これらはそれぞれ「json エンコードの内容をメモリに格納する」「メモリの中身を json デコードする」という関数

## encoding/json でストリームを扱う

### json.Decoder 型

encoding/json パッケージの中にある json.Decoder 型 29 は、ストリームから取得できるデータ
を json デコードする

### json.Encoder 型

encoding/json パッケージの中にある json.Encoder 型 30 は、json デコードの結果をストリーム
に流す

# aql ファイルを流すコマンド

mysql -h 127.0.0.1 -u user go_api_mysql -p < insertData.sql

# SQL ドライバー

## mySQL

db.Query メソッドの第二引数以降に、?に代入したい値を指定

## PostgresSQL

$1,$2

# memo

• POST /article: リクエストボディで受け取った記事を投稿する
– 構造体 models.Article を受け取って、それをデータベースに挿入する処理が必要
• GET /article/list: クエリパラメータ page で指定されたページ (1 ページに 5 個の記事
を表示) に表示するための記事一覧を取得する
– 指定された記事データをデータベースから取得して、それを models.Article 構造体
のスライス []models.Article に詰めて返す処理が必要
• GET /article/{id}: 指定 ID の記事を取得する
– 指定 ID の記事データをデータベースから取得して、それを models.Article 構造体の
形で返す処理が必要
– 指 定 ID の 記 事 に つ い た コ メ ン ト 一 覧 を デ ー タ ベ ー ス か ら 取 得 し て、 そ れ を
models.Comment 構造体のスライス []models.Comment に詰めて返す処理が必要
• POST /article/nice: 記事にいいねをつける
– 指定された ID の記事のいいね数を+1 するようにデータベースの中身を更新する処理
が必要
• POST /comment: リクエストボディで受け取ったコメントを投稿する
– 構造体 models.Comment を受け取って、それをデータベースに挿入する処理が必要

# sql.DB 型の　 Open/Close

sql.DB 型の頻繁な Open/Close は Go 公式としての非推奨事項
https://pkg.go.dev/database/sql#DB
