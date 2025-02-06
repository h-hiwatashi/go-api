# go-api

`APIを作りながら進むGo中級者への道`で API を作成するリポジトリ

## API エンドポイント一覧

| エンドポイント     | 概要                     | レスポンスに含めたい内容     |
| ------------------ | ------------------------ | ---------------------------- |
| POST /article      | 記事を投稿する           | 投稿に成功した記事の内容     |
| GET /article/list  | 記事一覧を取得する       | 記事のリスト                 |
| GET /article/{id}  | 指定 ID の記事を取得する | 記事の内容                   |
| POST /article/nice | 記事にいいねをつける     | いいねをつけた記事の内容     |
| POST /comment      | コメントを投稿する       | 投稿に成功したコメントの内容 |

## ディレクトリ構成

ハンドラ層
サービス層
レポジトリ層
