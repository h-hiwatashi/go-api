package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/h-hiwatashi/go-api/app/handlers"
	"github.com/h-hiwatashi/go-api/app/models"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/hello", handlers.HelloHandler).Methods(http.MethodGet)
	r.HandleFunc("/article", handlers.PostArticleHandler).Methods(http.MethodPost)
	r.HandleFunc("/article/list", handlers.ArticleListHandler).Methods(http.MethodGet)
	r.HandleFunc("/article/{id:[0-9]+}", handlers.ArticleDetailHandler).Methods(http.MethodGet)
	r.HandleFunc("/article/nice", handlers.PostNiceHandler).Methods(http.MethodPost)
	r.HandleFunc("/comment", handlers.PostCommentHandler).Methods(http.MethodPost)

	/// 標準パッケージ net/http だけで実装
	// http.HandleFunc("/hello", handlers.HelloHandler)
	// http.HandleFunc("/article", handlers.PostArticleHandler)
	// http.HandleFunc("/article/list", handlers.ArticleListHandler)
	// http.HandleFunc("/article/1", handlers.ArticleDetailHandler)
	// http.HandleFunc("/article/nice", handlers.PostNiceHandler)
	// http.HandleFunc("/comment", handlers.PostCommentHandler)

	// サーバー起動時のログを出力
	log.Println("server start at port 8080")

	// DBの接続
	dbUser := "user"
	dbPassword := "user"
	dbDatabase := "go_api_mysql"
	dbConn := fmt.Sprintf("%s:%s@tcp(go_api_mysql:3306)/%s?parseTime=true", dbUser,
		dbPassword, dbDatabase)
	db, err := sql.Open("mysql", dbConn)
	if err != nil {
		fmt.Println(err)
	}
	defer db.Close()

	// SQL実行
	articleID := 1
	const sqlStr = `
	select * 
	from articles
	where article_id = ?
	;`
	row := db.QueryRow(sqlStr, articleID)
	if err := row.Err(); err != nil {
		// データ取得件数が 0 件だった場合は
		// データ読み出し処理には移らずに終了
		fmt.Println(err)
		return
	}

	// データの格納
	// articleArray := make([]models.Article, 0)
	var article models.Article
	var createdTime sql.NullTime
	// 引数に「データ読み出し結果を格納したい変数のポインタ」を指定することで、rows の中に格納されている取得レコード内容を読み出す
	err = row.Scan(&article.ID, &article.Title, &article.Contents, &article.UserName, &article.NiceNum, &createdTime)
	if err != nil {
		fmt.Println(err)
		return
	}
	if createdTime.Valid {
		article.CreatedAt = createdTime.Time
	}
	fmt.Printf("%+v\n", articleArray)



	// ListenAndServe 関数にて、サーバーを起動
	// log.Fatal(http.ListenAndServe(":8080", nil))
	log.Fatal(http.ListenAndServe(":8080", r))
}
