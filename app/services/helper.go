package services

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

var (
	dbUser     = os.Getenv("MYSQL_DATABASE_USER")
	dbPassword = os.Getenv("MYSQL_DATABASE_PASSWORD")
	dbDatabase = os.Getenv("MYSQL_DATABASE")
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
