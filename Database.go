package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

type Author struct {
	Id int
	Name string
}
func main() {
	fmt.Println("Start App...")
	db, err := sql.Open("mysql", "root:root@tcp(localhost:3308)/go-example")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()
	insertAuthor("Anh Vuong Vu", db)
	fmt.Println(findByName("anh", db))


}

func insertAuthor(name string, db *sql.DB)  {
	insert, err := db.Query("insert into author (name) value (?)", name)
	if err != nil {
		panic(err.Error())
	}
	defer insert.Close()
}

func findByName(name string, db *sql.DB) []Author {
	authors := []Author{}
	result, err := db.Query("select * from author where name like ?", "%"+name+"%")
	if err != nil {
		panic(err.Error())
	}
	for result.Next() {
		var author Author
		err = result.Scan(&author.Id, &author.Name)
		if err != nil {
			panic(err.Error())
		}
		authors = append(authors, author)

	}
	return authors
}