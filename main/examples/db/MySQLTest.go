package main

import (
	"database/sql"
	_"github.com/go-sql-driver/mysql"
	"log"
)


func main() {

	db, err := sql.Open("mysql", "root:hechuan10131373@tcp(127.0.0.1:3306)/test")
	if err != nil {
		log.Fatal(err)
		return
	}

	var id int
	var name string

	rows, err := db.Query("select id, name from user;")
	if err != nil {
		log.Fatal(err)
		return
	}
	defer rows.Close()

	for rows.Next()  {
		rows.Scan(&id, &name)
		log.Println(id, name)
	}

	// 遍历过程中是否需到错误
	err = rows.Err()
	if err != nil {log.Fatal(err)}

	// 也可以使用statement来执行
	stmt, err := db.Prepare("INSERT INTO user VALUES (?, ?, ?, ?, ?);")
	if err != nil {log.Fatal(err)}
	res, err := stmt.Exec(4, "account4", "陈七", 23, 50)
	if err != nil {log.Fatal(err)}
	log.Println(res.LastInsertId())

	stmt, err = db.Prepare("SELECT id, name FROM user WHERE id = 4")
	rows, err = stmt.Query()
	for rows.Next()  {
		rows.Scan(&id, &name)
		log.Println(id, name)
	}

	defer db.Close();
}
