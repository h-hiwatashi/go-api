package services

import "github.com/h-hiwatashi/go-api/app/models"

// /article 関連を引き受けるサービス
type ArticleServicer interface {
	PostArticleService(article models.Article) (models.Article, error)
	GetArticleListService(page int) ([]models.Article, error)
	GetArticleService(articleID int) (models.Article, error)
	PostNiceService(article models.Article) (models.Article, error)
}

// /comment を引き受けるサービス
type CommentServicer interface {
	PostCommentService(comment models.Comment) (models.Comment, error)
}
