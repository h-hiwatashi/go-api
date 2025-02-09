package services

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

var (
	dbUser     = "user"
	dbPassword = "user"
	dbDatabase = "go_api_mysql"
	dbConn     = fmt.Sprintf(
		"%s:%s@tcp(localhost:3306)/%s?parseTime=true",
		dbUser,
		dbPassword,
		dbDatabase,
	)
)

func connectDB() (*sql.DB, error) {
	db, err := sql.Open("mysql", dbConn)
	if err != nil {
		return nil, err
	}
	return db, nil
}
