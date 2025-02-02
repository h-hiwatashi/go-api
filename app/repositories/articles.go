package repositories

import (
	"database/sql"

	"github.com/h-hiwatashi/go-api/app/models"
)

// 新規投稿をデータベースに insert する関数
// -> データベースに保存した記事内容と、発生したエラーを返り値にする
func InsertArticle(db *sql.DB, article models.Article) (models.Article, error) {
	const sqlStr = `
	insert into articles (title, contents, username, nice, created_at) values
	(?, ?, ?, 0, now());
	`
	var newArticle models.Article
	newArticle.Title, newArticle.Contents, newArticle.UserName = article.Title, article.Contents, article.UserName
	result, err := db.Exec(sqlStr, article.Title, article.Contents, article.UserName)
	if err != nil {
		return models.Article{}, err
	}
	id, _ := result.LastInsertId()

	newArticle.ID = int(id)

	return newArticle, nil
}

// 変数 page で指定されたページに表示する投稿一覧をデータベースから取得する関数
// -> 取得した記事データと、発生したエラーを返り値にする
func SelectArticleList(db *sql.DB, page int) ([]models.Article, error) {
	const sqlStr = `
	select article_id, title, contents, username, nice
	from articles
	limit ? offset ?;
	`
	offsetNum := (page - 1) * 5

	rows, err := db.Query(sqlStr, 5, offsetNum)
	if err != nil {
		return nil, err
	}
	articleArray := make([]models.Article, 0)

	for rows.Next() {
		var article models.Article
		err := rows.Scan(&article.ID, &article.Title, &article.Contents, &article.UserName, &article.NiceNum)
		if err != nil {
			return nil, err
		} else {
			articleArray = append(articleArray, article)
		}
	}
	return articleArray, nil
}

// 投稿 ID を指定して、記事データを取得する関数
// -> 取得した記事データと、発生したエラーを返り値にする
func SelectArticleDetail(db *sql.DB, articleID int) (models.Article, error) {
	const sqlStr = `  select *  from articles  where article_id =?;  `
	row := db.QueryRow(sqlStr, articleID)
	if err := row.Err(); err != nil {
		return models.Article{}, err
	}
	var article models.Article
	err := row.Scan(&article.ID, &article.Title, &article.Contents, &article.UserName, &article.NiceNum, &article.CreatedAt)
	if err != nil {
		return models.Article{}, err
	}
	return article, nil
}

// いいねの数を update する関数
// -> 発生したエラーを返り値にする
func UpdateNiceNum(db *sql.DB, articleID int) error {
	const sqlGetNice = `
	select nice
	from articles
	where article_id = ?;
	`
	const sqlUpdateNice = `update articles set nice = ? where article_id = ?`

	tx, err := db.Begin()
	if err != nil {
		return err
	}

	row := tx.QueryRow(sqlGetNice, articleID)
	if err := row.Err(); err != nil {
		tx.Rollback()
		return err
	}

	var nicenum int
	err = row.Scan(&nicenum)
	if err != nil {
		tx.Rollback()
		return err
	}

	_, err = tx.Exec(sqlUpdateNice, nicenum+1, articleID)
	if err != nil {
		tx.Rollback()
		return err
	}

	tx.Commit()
	return nil
}
