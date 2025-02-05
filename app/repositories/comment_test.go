package repositories_test

import (
	"testing"

	"github.com/h-hiwatashi/go-api/app/models"
	"github.com/h-hiwatashi/go-api/app/repositories"

	_ "github.com/go-sql-driver/mysql"
)

func TestSelectCommentList(t *testing.T) {
	articleID := 1
	got, err := repositories.SelectCommentList(testDB, articleID)
	if err != nil {
		t.Fatal(err)
	}

	for _, comment := range got {
		if comment.ArticleID != articleID {
			t.Errorf("want comment of articleID %d but got ID %d\n", articleID, comment.ArticleID)
		}
	}
}

func TestInsertComment(t *testing.T) {
	comment := models.Comment{
		ArticleID: 1,
		Message:   "test comment",
	}
	expectedCommentList := 5
	got, err := repositories.InsertComment(testDB, comment)
	if err != nil {
		t.Fatal(err)
	}

	articleID := 1
	commentList, err := repositories.SelectCommentList(testDB, articleID)
	if err != nil {
		t.Fatal(err)
	}

	if len(commentList) != expectedCommentList {
		t.Errorf("want comment of commentList length %d but got length %d\n", len(commentList), expectedCommentList)
	}

	t.Cleanup(func() {
		const sqlStr = `
			delete from comments
			where id =?;
		`
		testDB.Exec(sqlStr, got.CommentID)
	})
}
