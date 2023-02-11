package user

import (
	"database/sql"
	"fmt"
)


type User struct {
	id    int    `db:"id"`
	title string `db:"name"`
	data  string `db:"data"`
}

func ReadAll(db *sql.DB) []User{
	var users []User
	rows, err := db.Query("select * from users;")
	if err != nil {
		panic(err)
	}
	for rows.Next() {
		user := User{}
		err = rows.Scan(&user.id, &user.title, &user.data)
		if err != nil {
			panic(err)
		}
		users = append(users, user)
	}
	
	rows.Close()


	return users
}


func InsertData(db *sql.DB, name string, data string) {

	fmt.Println(data)

	res, err := db.Exec("INSERT INTO users (id, name, data) VALUES (0, ?, ?)",
		name,
		data,
	)
	fmt.Println(res)
	if err != nil {
		panic(err)
	}
	
}