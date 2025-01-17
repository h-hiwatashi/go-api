package handlers

import (
	"io"
	"log"
	"net/http"
)

func main() {
	// 他のパッケージからも参照可能な関数・変数・定数を作成するためには、その名前を大文字から始める必要があります
	func HelloHandler(w http.ResponseWriter, req *http.Request) {
		io.WriteString(w, "Hello, world!\n")
	}
	func PostArticleHandler(w http.ResponseWriter, req *http.Request) {
		io.WriteString(w, "Posting Article...\n")
	}
	func ArticleListHandler(w http.ResponseWriter, req *http.Request) {
		io.WriteString(w, "Article List\n")
	}
	func ArticleDetailHandler(w http.ResponseWriter, req *http.Request) {
		io.WriteString(w, "Article No.1\n")
	}
	func PostNiceHandler(w http.ResponseWriter, req *http.Request) {
		io.WriteString(w, "Posting Nice…\n")
	}
	func PostCommentHandler(w http.ResponseWriter, req *http.Request) {
		io.WriteString(w, "Posting Comment…\n")
	}

	// 定義した helloHandler を使うように登録
	http.HandleFunc("/", helloHandler)
	http.HandleFunc("/article", postArticleHandler)
	http.HandleFunc("/article/lis", articleListHandler)
	http.HandleFunc("/article/1", articleDetailHandler)
	http.HandleFunc("/article/nice", postNiceHandler)
	http.HandleFunc("/comment", postCommentHandler)

	// サーバー起動時のログを出力
	log.Println("server start at port 8080")

	// ListenAndServe 関数にて、サーバーを起動
	log.Fatal(http.ListenAndServe(":8080", nil))
}
