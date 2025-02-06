package repositories_test

import (
	"database/sql"
	"fmt"
	"os"
	"os/exec"
	"testing"

	_ "github.com/go-sql-driver/mysql"
)

var testDB *sql.DB

// DB 接続部分もついでに別関数に切り出し
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

func connectDB() error {
	var err error
	testDB, err = sql.Open("mysql", dbConn)
	if err != nil {
		return err
	}
	return nil
}

func setup() error {
	// DBの接続
	if err := connectDB(); err != nil {
		return err
	}
	if err := cleanupDB(); err != nil {
		fmt.Println("cleanup", err)
		return err
	}
	if err := setupTestData(); err != nil {
		fmt.Println("setup")
		return err
	}
	return nil
}

// テスト共通の後処理
func teardown() {
	cleanupDB()
	testDB.Close()
}

func TestMain(m *testing.M) {
	err := setup()
	if err != nil {
		os.Exit(1)
	}

	// m.Run()とはテストパッケージ内にある全てのテストを実行する関数
	m.Run()

	teardown()
}

// mysql コマンドを用いて、setupDB.sql と cleanupDB.sql ファイルに書いたクエリを実行
func setupTestData() error {
	cmd := exec.Command("mysql", "-h", "127.0.0.1", "-u", "user", "go_api_mysql",
		"--password=user", "-e", "source ./testdata/setupDB.sql")
	err := cmd.Run()
	if err != nil {
		return err
	}
	return nil
}

func cleanupDB() error {
	cmd := exec.Command("mysql", "-h", "127.0.0.1", "-u", "user", "go_api_mysql",
		"--password=user", "-e", "source ./testdata/cleanupDB.sql")
	err := cmd.Run()
	if err != nil {
		return err
	}
	return nil
}
