package repositories

import (
	"database/sql"

	"github.com/h-hiwatashi/go-api/app/models"
)

// 新規投稿をデータベースに insert する関数
// -> データベースに保存したコメント内容と、発生したエラーを返り値にする
func InsertComment(db *sql.DB, comment models.Comment) (models.Comment, error) {
	const sqlStr = `
	insert into comments (article_id, message, created_at) values
	(?, ?, now());
	`
	// (問 5) 構造体 models.Comment を受け取って、それをデータベースに挿入する処理
	var newComment models.Comment
	result, err := db.Exec(sqlStr, comment.ArticleID, comment.Message)
	if err != nil {
		return models.Comment{}, err
	}
	id, _ := result.LastInsertId()

	newComment.CommentID = int(id)

	return newComment, nil
}

// 指定 ID の記事についたコメント一覧を取得する関数
// -> 取得したコメントデータと、発生したエラーを返り値にする
func SelectCommentList(db *sql.DB, articleID int) ([]models.Comment, error) {
	const sqlStr = `
	select *
	from comments
	where article_id = ?;
	`
	// (問 6) 指定 ID の記事についたコメント一覧をデータベースから取得して、
	// それを`models.Comment`構造体のスライス`[]models.Comment`に詰めて返す処理
	rows, err := db.Query(sqlStr, articleID)
	if err != nil {
		return nil, err
	}
	commentArray := make([]models.Comment, 0)

	for rows.Next() {
		var comment models.Comment
		err := rows.Scan(&comment.CommentID, &comment.ArticleID, &comment.Message, &comment.CreatedAt)
		if err != nil {
			return nil, err
		} else {
			commentArray = append(commentArray, comment)
		}
	}
	return commentArray, nil
}
