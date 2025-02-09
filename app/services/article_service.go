package services

import (
	"github.com/h-hiwatashi/go-api/app/models"
	"github.com/h-hiwatashi/go-api/app/repositories"
)

// 指定IDの記事をDBから取得する関数
func GetArticleService(articleID int) (models.Article, error) {
	db, err := connectDB()
	if err != nil {
		return models.Article{}, err
	}
	defer db.Close()

	article, err := repositories.SelectArticleDetail(db, articleID)
	if err != nil {
		return models.Article{}, err
	}

	// SelectArticleDetail 関数では「指定 ID 記事に紐づいたコメント一覧」までは取得できないため、
	// SelectCommentList 関数を実行する
	commentList, err := repositories.SelectCommentList(db, articleID)
	if err != nil {
		return models.Article{}, err
	}

	article.CommentList = append(article.CommentList, commentList...)
	return article, nil
}
