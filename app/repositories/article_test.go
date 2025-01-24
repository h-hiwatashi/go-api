package repositories_test

import (
	"database/sql"
	"fmt"
	"testing"

	"github.com/h-hiwatashi/go-api/app/models"
	"github.com/h-hiwatashi/go-api/app/repositories"

	_ "github.com/go-sql-driver/mysql"
)

func TestSelectArticleDetail(t *testing.T) {
	// DBの接続
	dbUser := "user"
	dbPassword := "user"
	dbDatabase := "go_api_mysql"
	dbConn := fmt.Sprintf("%s:%s@tcp(go_api_mysql:3306)/%s?parseTime=true", dbUser,
		dbPassword, dbDatabase)
	db, err := sql.Open("mysql", dbConn)
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()

	expected := models.Article{
		ID:       1,
		Title:    "firstPost",
		Contents: "This is my first blog",
		UserName: "user name",
		NiceNum:  2,
	}
	got, err := repositories.SelectArticleDetail(db, expected.ID)

	if err != nil {
		t.Fatal(err)
	}
	if got.ID != expected.ID {
		t.Errorf("ID: get %d but want %d\n", got.ID, expected.ID)
	}
	if got.Title != expected.Title {
		t.Errorf("Title: get %s but want %s\n", got.Title, expected.Title)
	}
	if got.Contents != expected.Contents {
		t.Errorf("Content: get %s but want %s\n", got.Contents, expected.Contents)
	}
	if got.UserName != expected.UserName {
		t.Errorf("UserName: get %s but want %s\n", got.UserName, expected.UserName)
	}
	if got.NiceNum != expected.NiceNum {
		t.Errorf("NiceNum: get %d but want %d\n", got.NiceNum, expected.NiceNum)
	}
}
