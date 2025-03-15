package controllers_test

import (
	// (略)

	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

func TestArticleListHandler(t *testing.T) {
	// 2. テスト対象のハンドラメソッドに入れる input を定義
	var tests = []struct {
		name       string
		query      string
		resultCode int
	}{
		{name: "number query", query: "1", resultCode: http.StatusOK},
		{name: "alphabet query", query: "aaa", resultCode: http.StatusBadRequest},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// ハンドラに渡す 2 つの引数
			// w http.ResponseWriter, req *http.Request を用意する
			url := fmt.Sprintf("http://localhost:8080/article/list?page=%s", tt.query)
			// 引数: method, target string, body io.Reader
			req := httptest.NewRequest(http.MethodGet, url, nil)
			res := httptest.NewRecorder()

			// 3. テスト対象のハンドラメソッドから output を得る
			aCon.ArticleListHandler(res, req)
			// 4. output が期待通りかチェック
			if res.Code != tt.resultCode {
				t.Errorf("unexpected StatusCode: want %d but %d\n", tt.resultCode, res.Code)
			}

		},
		)
	}
}

func TestArticleDetailHandler(t *testing.T) {
	var tests = []struct {
		name       string
		articleID  string
		resultCode int
	}{
		{name: "number pathparam", articleID: "1", resultCode: http.StatusOK},
		{name: "alphabet pathparam", articleID: "aaa", resultCode: http.StatusNotFound},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			url := fmt.Sprintf("http://localhost:8080/article/%s", tt.articleID)
			req := httptest.NewRequest(http.MethodGet, url, nil)

			res := httptest.NewRecorder()

			// gorilla/mux のルータを用意
			r := mux.NewRouter()
			// パス/article/{id}とハンドラメソッド ArticleDetailHandler の対応付けを登録
			r.HandleFunc("/article/{id:[0-9]+}", aCon.ArticleDetailHandler).Methods(http.MethodGet)
			r.ServeHTTP(res, req)

			if res.Code != tt.resultCode {
				t.Errorf("unexpected StatusCode: want %d but %d\n", tt.resultCode, res.Code)
			}
		})
	}
}
