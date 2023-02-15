package repositories

import (
	"database/sql"

	"github.com/ninomae42/go_api_book/models"
)

func SelectCommentList(db *sql.DB, articleID int) ([]models.Comment, error) {
	const sqlStr = `
		select
			*
		from
			comments
		where article_id = ?;
	`

	rows, err := db.Query(sqlStr, articleID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	commentArray := make([]models.Comment, 0)
	for rows.Next() {
		var comment models.Comment
		var createdTime sql.NullTime
		rows.Scan(&comment.CommentID, &comment.ArticleID, &comment.Message, &createdTime)
		if createdTime.Valid {
			comment.CreatedAt = createdTime.Time
		}
		commentArray = append(commentArray, comment)
	}

	return commentArray, nil
}
