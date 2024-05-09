package main

import (
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

type Item struct {
	Id   int64  `db:"id"`
	Name string `db:"name"`
}

func main() {
	sqldb, err := sqlx.Open("mysql", "root:root@tcp(localhost:13306)/prjctr")
	if err != nil {
		log.Fatal(err)
	}

	sqldb.MustExec(`
		CREATE TABLE IF NOT EXISTS items (
		id   BIGINT NOT NULL AUTO_INCREMENT,
		name varchar(250) NOT NULL,
		PRIMARY KEY (id));
	`)

	newItem := Item{Name: "New Item 4"}
	_, err = sqldb.NamedExec("INSERT INTO items (name) VALUES (:name)", &newItem)
	if err != nil {
		log.Fatal(err)
	}

	items := []*Item{}
	err = sqldb.Select(&items, "SELECT * FROM items ORDER BY id ASC")
	if err != nil {
		log.Fatal(err)
	}
	for _, item := range items {
		fmt.Println(*item)
	}

	err = sqldb.Select(&items, "SELECT * FROM items WHERE name = ?", "New Item 4")
	if err != nil {
		log.Fatal(err)
	}
	for _, item := range items {
		fmt.Println(*item)
	}
}
