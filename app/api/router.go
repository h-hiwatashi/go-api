package api

import (
	"database/sql"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/h-hiwatashi/go-api/app/controllers"
	"github.com/h-hiwatashi/go-api/app/services"
)

func NewRouter(db *sql.DB) *mux.Router {
	// 2. sql.DB 型をもとに、サーバー全体で使用するサービス構造体 MyAppService を一つ生成する
	ser := services.NewMyAppService(db)
	// 3. MyAppService 型をもとに、サーバー全体で使用するコントローラ構造体 MyAppControllerを一つ生成する
	aCon := controllers.NewArticleController(ser)
	cCon := controllers.NewCommentController(ser)

	r := mux.NewRouter()
	r.HandleFunc("/hello", aCon.HelloHandler).Methods(http.MethodGet)
	r.HandleFunc("/article", aCon.PostArticleHandler).Methods(http.MethodPost)
	r.HandleFunc("/article/list", aCon.ArticleListHandler).Methods(http.MethodGet)
	r.HandleFunc("/article/{id:[0-9]}", aCon.ArticleDetailHandler).Methods(http.MethodGet)
	r.HandleFunc("/article/nice", aCon.PostNiceHandler).Methods(http.MethodPost)
	r.HandleFunc("/comment", cCon.PostCommentHandler).Methods(http.MethodPost)

	return r
}
