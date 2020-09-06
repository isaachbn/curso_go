package main

import (
	"database/sql"
	_"github.com/go-sql-driver/mysql"
	"log"
)

func main() {
	db, err := sql.Open("mysql", "root:root@/cursogo")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	tx, _ := db.Begin()
	stmt, _ := tx.Prepare("insert into usuarios(id, nome) values(?,?)")
	stmt.Exec(100, "Isaac Henrique")
	stmt.Exec(1001, "Lorena Farias")
	_, err = stmt.Exec(1, "Thiago") //Chave duplicada

	if err != nil {
		tx.Rollback()
		log.Fatal(err)
	}

	tx.Commit()
}