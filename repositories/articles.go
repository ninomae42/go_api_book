package repositories

import (
	"database/sql"

	"github.com/ninomae42/go_api_book/models"
)

// 新規投稿をDBにinsertする関数
func InsertArticle(db *sql.DB, article models.Article) (models.Article, error) {
	const sqlStr = `insert into articles (title, contents, username, nice, created_at) values
		(?, ?, ?, 0, now());
	`

	var newArticle models.Article
	newArticle.Title = article.Title
	newArticle.Contents = article.Contents
	newArticle.UserName = article.UserName
	
	result, err := db.Exec(sqlStr, article.Title, article.Contents, article.UserName)
	if err != nil {
		return models.Article{}, err
	}

	id, _ := result.LastInsertId()
	newArticle.ID = int(id)
	return newArticle, nil
}
