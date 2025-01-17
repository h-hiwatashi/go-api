package main

import (
	"log"
	"net/http"
	"github.com/h-hiwatashi/go-api/app/handlers"
)

func main() {
	http.HandleFunc("/hello", handlers.HelloHandler)
	http.HandleFunc("/article", handlers.PostArticleHandler)
	http.HandleFunc("/article/list", handlers.ArticleListHandler)
	http.HandleFunc("/article/1", handlers.ArticleDetailHandler)
	http.HandleFunc("/article/nice", handlers.PostNiceHandler)
	http.HandleFunc("/comment", handlers.PostCommentHandler)

	// サーバー起動時のログを出力
	log.Println("server start at port 8080")

	// ListenAndServe 関数にて、サーバーを起動
	log.Fatal(http.ListenAndServe(":8080", nil))
}
