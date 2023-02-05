package main

import (
	"database/sql" //データベースに接続する時に使うパッケージ
	"fmt"
	"go_docker/article"
	"log"
	"os"
	"time"
	
	_ "github.com/go-sql-driver/mysql" // 今回はMySQLに接続する  パッケージ内のinitメソッドだけが実行
)

func open(path string, count uint) *sql.DB {
	db, err := sql.Open("mysql", path)
	if err != nil {
		log.Fatal("open error:", err)
	}

	// db.Ping()でdbサービスへの接続確認をし、失敗したらcountを減らし、openメソッドを再起で呼びます
	if err = db.Ping(); err != nil {
		time.Sleep(time.Second * 2)
		count--
		fmt.Printf("retry... count:%v\n", count)
		return open(path, count)
	}

	fmt.Println("db connected!!")
	return db
}

func connectDB() *sql.DB {
	// database/sqlで指定する接続情報
	// [user名]:[password]@tcp([host名]:[port])/database名
	// host名はdb？
	var path string = fmt.Sprintf("%s:%s@tcp(db:3306)/%s?charset=utf8&parseTime=true",
		os.Getenv("MYSQL_USER"), os.Getenv("MYSQL_PASSWORD"),
		os.Getenv("MYSQL_DATABASE"))

	return open(path, 100)
}

func main() {
	db := connectDB()
	defer db.Close()// article.ReadAll(db)後に実行
	article.ReadAll(db)
}