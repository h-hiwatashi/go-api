package services

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

var (
	// TODO: 環境変数から取得するように変更
	// dbUser     = os.Getenv("MYSQL_DATABASE_USER")
	// dbPassword = os.Getenv("MYSQL_DATABASE_PASSWORD")
	// dbDatabase = os.Getenv("MYSQL_DATABASE")
	dbUser     = "user"
	dbPassword = "user"
	dbDatabase = "go_api_mysql"
	dbConn     = fmt.Sprintf(
		"%s:%s@tcp(go_api_mysql:3306)/%s?parseTime=true",
		dbUser,
		dbPassword,
		dbDatabase,
	)
)

func connectDB() (*sql.DB, error) {
	fmt.Println(dbConn)
	db, err := sql.Open("mysql", dbConn)
	if err != nil {
		return nil, err
	}
	return db, nil
}
