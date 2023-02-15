package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/ninomae42/go_api_book/models"
)

func main() {
	dbUser := "docker"
	dbPassword := "docker"
	dbDatabase := "sampledb"
	dbConn := fmt.Sprintf("%s:%s@tcp(127.0.0.1:3306)/%s?parseTime=True", dbUser, dbPassword, dbDatabase)

	db, err := sql.Open("mysql", dbConn)
	if err != nil {
		fmt.Println(err)
	}
	defer db.Close()

	articleID := 1
	const sqlStr = `
		select *
		from articles
		where article_id = ?;
	`

	var article models.Article
	var createdTime sql.NullTime
	row := db.QueryRow(sqlStr, articleID)
	if err := row.Err(); err != nil {
		fmt.Println(err)
		return
	}

	err = row.Scan(&article.ID, &article.Title, &article.Contents, &article.UserName,
		&article.NiceNum, &createdTime)
	if err != nil {
		fmt.Println(err)
		return
	}
	if createdTime.Valid {
		article.CreatedAt = createdTime.Time
	}
	fmt.Printf("%+v\n", article)
}
