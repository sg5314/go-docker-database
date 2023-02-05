package article

import (
	"database/sql"
	"fmt"
	"log"
)

type Article struct {
	id    int
	title string
	body  string
}

func ReadAll(db *sql.DB) {
	var articles []Article
	rows, err := db.Query("select * from article;")
	if err != nil {
		log.Fatalf("%v", err)
		panic(err) // ランタイムエラー  recover？
	}
	for rows.Next() {
		article := Article{}//初期化
		// 引数に渡したポインタにレコードの内容を読み込み
		err = rows.Scan(&article.id, &article.title, &article.body)
		if err != nil {
			log.Fatalf("%v", err)
			panic(err) // ランタイムエラー
		}
		articles = append(articles, article)
	}
	rows.Close()

	fmt.Println(articles)
}


// SELECT等を利用：Query
// 単一データを取得：QueryRow
// INSERT, CREATE, UPDATE, DELETE等を利用：Exec