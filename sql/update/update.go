package main

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

func main() {
	db, err := sql.Open("mysql", "root:root@/cursogo")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	tx, _ := db.Begin()
	stmt, _ := tx.Prepare("update usuarios set nome = ? where id = ?")
	_, err = stmt.Exec("Isaac Henrique", 1)

	if err != nil {
		tx.Rollback()
		log.Fatal(err)
	}

	stmt2, _ := tx.Prepare("delete from usuarios where id = ?")
	_, err = stmt2.Exec(3)

	if err != nil {
		tx.Rollback()
		log.Fatal(err)
	}

	tx.Commit()
}
