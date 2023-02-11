package main

import (
	// "github.com/gin-gonic/gin"
	"database/sql"// golangでデータベースに接続する時に使うパッケージ
	"log"
	"os"
	"fmt"
	"time"
	"server/user"


	_ "github.com/go-sql-driver/mysql"// MySQLに接続するのでドライバ  パッケージの前に_(アンダースコア)を付けるとパッケージ内のinitメソッドだけが実行
)

func open(path string, count uint) *sql.DB {
	db, err := sql.Open("mysql", path)
	if err != nil {
		log.Fatal("open error:", err)
	}

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
	var path string = fmt.Sprintf("%s:%s@tcp(db:3306)/%s?charset=utf8&parseTime=true",
		os.Getenv("MYSQL_USER"), os.Getenv("MYSQL_PASSWORD"),
		os.Getenv("MYSQL_DATABASE"))

	return open(path, 100)
}


// https://zenn.dev/hsaki/articles/go-convert-json-struct#go%E6%A7%8B%E9%80%A0%E4%BD%93%E3%81%8B%E3%82%89json%E3%81%B8%E3%81%AE%E3%82%A8%E3%83%B3%E3%82%B3%E3%83%BC%E3%83%89


func CreateData(){
	for 
}

func main() {

	db := connectDB()
	defer db.Close()

	var get_data string
	get_data := CreateData()

	user.InsertData(db, get_data)

	var datas []user.User
	datas = user.ReadAll(db)
	

	fmt.Println(datas)

	// r := gin.Default()
	// r.GET("/ping", func(c *gin.Context) {
	// 	c.JSON(200, gin.H{
	// 		"message": "pong",
	// 	})
	// })
	// r.Run(":3000") // 0.0.0.0:8080 でサーバーを立てます。
}