# このディレクトリについて

- データベースと通信して、データの取得・挿入を行う

## テスト

| repositories パッケージの関数 | 概要                                   | テストの実装場所 |
| ----------------------------- | -------------------------------------- | ---------------- |
| InsertArticle 関数            | ブログ記事をデータベースに insert      | article_test.go  |
| SelectArticleList 関数　      | ブログ記事一覧を取得                   | article_test.go  |
| SelectArticleDetail 関数      | 指定 ID の記事データを取得             | article_test.go  |
| UpdateNiceNum 関数            | 指定記事のいいね数を+1 する            | article_test.go  |
| SelectCommentList 関数        | 指定 ID 記事についたコメント一覧を取得 | comment_test.go  |
| InsertComment 関数            | コメントをデータベースに insert        | comment_test.go  |
