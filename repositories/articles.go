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

// 投稿IDを指定して、記事データを取得する関数
func SelectArticleDetail(db *sql.DB, articleID int) (models.Article, error) {
	const sqlStr = `
		select *
		from articles
		where article_id = ?;
	`

	row := db.QueryRow(sqlStr, articleID)
	if err := row.Err(); err != nil {
		return models.Article{}, err
	}

	var article models.Article
	var createdTime sql.NullTime
	err := row.Scan(&article.ID, &article.Title, &article.Contents, &article.UserName, &article.NiceNum, &createdTime)
	if err != nil {
		return models.Article{}, err
	}

	if createdTime.Valid {
		article.CreatedAt = createdTime.Time
	}

	return article, nil
}
