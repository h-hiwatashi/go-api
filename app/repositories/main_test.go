package repositories_test

import (
	"database/sql"
	"fmt"
	"os"
	"testing"

	_ "github.com/go-sql-driver/mysql"
)

var testDB *sql.DB

func setup() error {
	// DBの接続
	dbUser := "user"
	dbPassword := "user"
	dbDatabase := "go_api_mysql"
	dbConn := fmt.Sprintf("%s:%s@tcp(go_api_mysql:3306)/%s?parseTime=true", dbUser,
		dbPassword, dbDatabase)
	var err error
	testDB, err = sql.Open("mysql", dbConn)
	if err != nil {
		return err
	}
	return nil
}

func teardown() {
	testDB.Close()
}

func TestMain(m *testing.M) {
	err := setup()
	if err != nil {
		os.Exit(1)
	}
	m.Run()
	teardown()
}
