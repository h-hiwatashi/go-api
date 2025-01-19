package handlers

import (
	"encoding/json"
	"errors"
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

// POST /article のハンドラ
func PostArticleHandler(w http.ResponseWriter, req *http.Request) {
	// 1. バイトスライス reqBodybuffer を何らかの形で用意
	// req.Header の Get メソッドを呼ぶことで、リクエストヘッダの Content-Length フィールドの値を取得
	length, err := strconv.Atoi(req.Header.Get("Content-Length"))
	// int 型への変換に失敗した場合、400 番エラーを返却
	if err != nil {
		http.Error(w, "cannot get content length\n", http.StatusBadRequest)
		return
	}
	reqBodybuffer := make([]byte, length)
	// 2. Read メソッドでリクエストボディを読み出し
	// 戻り値 err に、読み取り時に起きたエラーの内容が格納される
	// errors.Is 関数は、第一引数として渡された err が第二引数 target として渡されたエラーと一致するかどうかを判定する関数
	if _, err := req.Body.Read(reqBodybuffer); !errors.Is(err, io.EOF) {
		// Read メソッドからの err が io.EOF 以外だった場合、500 番エラーを返却
		http.Error(w, "fail to get request body\n", http.StatusBadRequest)
		return
	}
	// 3. ボディを Close する
	defer req.Body.Close()

	//Article 型の変数 reqArticle の中に、 reqBodybuffer に格納された json バイト列をデコードした結果を格納
	var reqArticle models.Article
	if err := json.Unmarshal(reqBodybuffer, &reqArticle); err != nil {
		http.Error(w, "fail to decode json\n", http.StatusBadRequest)
		return
	}

	article := reqArticle
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
