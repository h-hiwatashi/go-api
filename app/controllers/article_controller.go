package controllers

import (
	"encoding/json"

	"io"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/h-hiwatashi/go-api/app/apperrors"
	"github.com/h-hiwatashi/go-api/app/models"

	"github.com/h-hiwatashi/go-api/app/controllers/services"
)

// 1. コントローラ構造体を定義
type ArticleController struct {
	// 2. フィールドに MyAppService 構造体を含める
	service services.ArticleServicer
}

// コンストラクタの定義
func NewArticleController(s services.ArticleServicer) *ArticleController {
	return &ArticleController{service: s}
}

// 他のパッケージからも参照可能な関数・変数・定数を作成するためには、その名前を大文字から始める必要があります
func (c *ArticleController) HelloHandler(w http.ResponseWriter, req *http.Request) {
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
func (c *ArticleController) ArticleListHandler(w http.ResponseWriter, req *http.Request) {
	queryMap := req.URL.Query()

	// クエリパラメータpageを取得
	var page int
	/// もし map 型の変数 queryMap が文字列"page"をキーに持っているのであれば、p には pageキーに対応する値 queryMap["page"] が、ok には true が格納される
	/// もし map 型の変数 queryMap が文字列"page"をキーに持っていないのであれば、ok にはfalse が格納される
	if p, ok := queryMap["page"]; ok && len(p) > 0 {
		var err error
		page, err = strconv.Atoi(p[0])
		if err != nil {
			err = apperrors.BadParam.Wrap(err, "queryparam must be number")
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
func (c *ArticleController) PostArticleHandler(w http.ResponseWriter, req *http.Request) {
	//Article 型の変数 reqArticle の中に、 reqBodybuffer に格納された json バイト列をデコードした結果を格納
	var reqArticle models.Article
	if err := json.NewDecoder(req.Body).Decode(&reqArticle); err != nil {
		// json デコードしたときに得られたエラーをラップして、エラーメッセージを出力
		err = apperrors.ReqBodyDecodeFailed.Wrap(err, "bad request body")
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
func (c *ArticleController) ArticleDetailHandler(w http.ResponseWriter, req *http.Request) {
	// articleID という変数に、リクエストの URL から取得した id パラメータを格納
	// strconv.Atoi で文字列を数値に変換
	articleID, err := strconv.Atoi(mux.Vars(req)["id"])
	if err != nil {
		err = apperrors.BadParam.Wrap(err, "queryparam must be number")
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
func (c *ArticleController) PostNiceHandler(w http.ResponseWriter, req *http.Request) {
	var reqArticle models.Article
	if err := json.NewDecoder(req.Body).Decode(&reqArticle); err != nil {
		err = apperrors.ReqBodyDecodeFailed.Wrap(err, "bad request body")
		http.Error(w, "fail to decode json\n", http.StatusBadRequest)
	}

	article, err := c.service.PostNiceService(reqArticle)
	if err != nil {
		http.Error(w, "fail internal exec\n", http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(article)
}
