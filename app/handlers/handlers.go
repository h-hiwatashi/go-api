package handlers

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/h-hiwatashi/go-api/app/models"
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
	article := models.Article1
	jsonData, err := json.Marshal(article)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	// []byte 型である jsonData を
	// string 型に変換することなくそのままレスポンス書き込みに使うことができる
	w.Write(jsonData)
}

// • x が数字だった場合は、記事一覧ページの x ページ目に表示されるデータを返す
// • page に対応する値が複数個送られてきた場合には、最初の値を使用する
// • x が数字ではなかった場合には、リクエストについていたパラメータの値が悪いということなので 400 番エラーを返す
// • クエリパラメータが URL についていなかった場合には、パラメータ page=1 がついていたときと同じ処理をする
func ArticleListHandler(w http.ResponseWriter, req *http.Request) {
	// queryMap := req.URL.Query()
	// var page int
	// /// もし map 型の変数 queryMap が文字列"page"をキーに持っているのであれば、p には pageキーに対応する値 queryMap["page"] が、ok には true が格納される
	// /// もし map 型の変数 queryMap が文字列"page"をキーに持っていないのであれば、ok にはfalse が格納される
	// if p, ok := queryMap["page"]; ok && len(p) > 0 {
	// 	var err error
	// 	page, err = strconv.Atoi(p[0])
	// 	if err != nil {
	// 		http.Error(w, "Invalid page", http.StatusBadRequest)
	// 		return
	// 	}
	// } else {
	// 	page = 1
	// }
	// reqString := fmt.Sprintf("Article List (page: %d)\n", page)

	articleList := []models.Article{models.Article1, models.Article2}
	jsonData, err := json.Marshal(articleList)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	w.Write(jsonData)
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

	article := models.Article1
	jsonData, err := json.Marshal(article)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	w.Write(jsonData)
}
func PostNiceHandler(w http.ResponseWriter, req *http.Request) {
	article := models.Article1
	jsonData, err := json.Marshal(article)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	w.Write(jsonData)
}
func PostCommentHandler(w http.ResponseWriter, req *http.Request) {
	comment := models.Comment1
	jsonData, err := json.Marshal(comment)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	w.Write(jsonData)
}
