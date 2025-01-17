package main

import (
	"io"
	"log"
	"net/http"
)

func main() {
	helloHandler := func(w http.ResponseWriter, req *http.Request) {
		io.WriteString(w, "Hello, world!\n")
	}
	// 定義した helloHandler を使うように登録
	http.HandleFunc("/", helloHandler)
	// サーバー起動時のログを出力
	log.Println("server start at port 8080")
	// ListenAndServe 関数にて、サーバーを起動
	log.Fatal(http.ListenAndServe(":8080", nil))
}
