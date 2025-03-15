package controllers_test

import (
	// (略)

	"testing"

	_ "github.com/go-sql-driver/mysql"
	"github.com/h-hiwatashi/go-api/app/controllers"
	"github.com/h-hiwatashi/go-api/app/controllers/testdata"
)

var aCon *controllers.ArticleController

func TestMain(m *testing.M) {
	ser := testdata.NewServiceMock()
	aCon = controllers.NewArticleController(ser)

	m.Run()
}
