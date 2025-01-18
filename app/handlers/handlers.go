package handlers

import (
	"fmt"
	"io"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// 他のパッケージからも参照可能な関数・変数・定数を作成するためには、その名前を大文字から始める必要があります
func HelloHandler(w http.ResponseWriter, req *http.Request) {
	// if req.Method == http.MethodGet {
	// 	io.WriteString(w, "Hello, world!\n")
	// } else {
	// 	// もし、req の中の Method フィールドが GET でなかったら
	// 	// Invalid method というレスポンスを、405 番のステータスコードと共に返す
	// 	http.Error(w, "Invalid method", http.StatusMethodNotAllowed)
	// }
	io.WriteString(w, "Hello, world!\n")
}
func PostArticleHandler(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "Posting Article...\n")
}
func ArticleListHandler(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "Article List\n")
}
func ArticleDetailHandler(w http.ResponseWriter, req *http.Request) {
	// articleID という変数に、リクエストの URL から取得した id パラメータを格納
	// strconv.Atoi で文字列を数値に変換
	articleID, err := strconv.Atoi(mux.Vars(req)["id"])
	if err != nil {
		// 400 番エラー (BadRequest) を返す
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}
	resString := fmt.Sprintf("Article No.%d\n", articleID)
	io.WriteString(w, resString)
}
func PostNiceHandler(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "Posting Nice…\n")
}
func PostCommentHandler(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "Posting Comment…\n")
}
