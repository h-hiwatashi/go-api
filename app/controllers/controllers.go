package controllers

import (
	"encoding/json"
	// "errors"

	"io"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/h-hiwatashi/go-api/app/models"

	// "github.com/h-hiwatashi/go-api/app/repositories"
	"github.com/h-hiwatashi/go-api/app/services"
)

// 1. コントローラ構造体を定義
type MyAppController struct {
	// 2. フィールドに MyAppService 構造体を含める
	service *services.MyAppService
}

// コンストラクタの定義
func NewMyAppController(s *services.MyAppService) *MyAppController {
	return &MyAppController{service: s}
}

// 他のパッケージからも参照可能な関数・変数・定数を作成するためには、その名前を大文字から始める必要があります
func (c *MyAppController) HelloHandler(w http.ResponseWriter, req *http.Request) {
	// if req.Method == http.MethodGet {
	// 	io.WriteString(w, "Hello, world!\n")
	// } else {
	// もし、req の中の Method フィールドが GET でなかったら
	// Invalid method というレスポンスを、405 番のステータスコードと共に返す
	// 	http.Error(w, "Invalid method", http.StatusMethodNotAllowed)
	// }
	io.WriteString(w, "Hello, world!\n")
}

// GET /article/list
// • x が数字だった場合は、記事一覧ページの x ページ目に表示されるデータを返す
// • page に対応する値が複数個送られてきた場合には、最初の値を使用する
// • x が数字ではなかった場合には、リクエストについていたパラメータの値が悪いということなので 400 番エラーを返す
// • クエリパラメータが URL についていなかった場合には、パラメータ page=1 がついていたときと同じ処理をする
func (c *MyAppController) ArticleListHandler(w http.ResponseWriter, req *http.Request) {
	queryMap := req.URL.Query()

	// クエリパラメータpageを取得
	var page int
	/// もし map 型の変数 queryMap が文字列"page"をキーに持っているのであれば、p には pageキーに対応する値 queryMap["page"] が、ok には true が格納される
	/// もし map 型の変数 queryMap が文字列"page"をキーに持っていないのであれば、ok にはfalse が格納される
	if p, ok := queryMap["page"]; ok && len(p) > 0 {
		var err error
		page, err = strconv.Atoi(p[0])
		if err != nil {
			http.Error(w, "Invalid query parameter", http.StatusBadRequest)
			return
		}
	} else {
		page = 1
	}

	articleList, err := c.service.GetArticleListService(page)
	if err != nil {
		http.Error(w, "fail internal exec\n", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(articleList)
}

// POST /article のハンドラ
func (c *MyAppController) PostArticleHandler(w http.ResponseWriter, req *http.Request) {
	//Article 型の変数 reqArticle の中に、 reqBodybuffer に格納された json バイト列をデコードした結果を格納
	var reqArticle models.Article
	if err := json.NewDecoder(req.Body).Decode(&reqArticle); err != nil {
		http.Error(w, "fail to decode json\n", http.StatusBadRequest)
	}

	article, err := c.service.PostArticleService(reqArticle)
	if err != nil {
		http.Error(w, "fail internal exec\n", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(article)
}

// GET /article/:id
func (c *MyAppController) ArticleDetailHandler(w http.ResponseWriter, req *http.Request) {
	// articleID という変数に、リクエストの URL から取得した id パラメータを格納
	// strconv.Atoi で文字列を数値に変換
	articleID, err := strconv.Atoi(mux.Vars(req)["id"])
	if err != nil {
		// 400 番エラー (BadRequest) を返す
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	article, err := c.service.GetArticleService(articleID)
	if err != nil {
		http.Error(w, "fail internal exec\n", http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(article)
}

// POST /article/nice の挙動確認用
func (c *MyAppController) PostNiceHandler(w http.ResponseWriter, req *http.Request) {
	var reqArticle models.Article
	if err := json.NewDecoder(req.Body).Decode(&reqArticle); err != nil {
		http.Error(w, "fail to decode json\n", http.StatusBadRequest)
	}

	article, err := c.service.PostNiceService(reqArticle)
	if err != nil {
		http.Error(w, "fail internal exec\n", http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(article)
}

// POST /comment
func (c *MyAppController) PostCommentHandler(w http.ResponseWriter, req *http.Request) {
	var reqComment models.Comment
	if err := json.NewDecoder(req.Body).Decode(&reqComment); err != nil {
		http.Error(w, "fail to decode json\n", http.StatusBadRequest)
	}
	comment, err := c.service.PostCommentService(reqComment)
	if err != nil {
		http.Error(w, "fail internal exec\n", http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(comment)
}
