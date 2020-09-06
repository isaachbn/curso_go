package main

import (
	"database/sql"
	json2 "encoding/json"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"net/http"
	"strconv"
	"strings"
)

// Usuario
type Usuario struct {
	ID   int    `json:"id"`
	Nome string `json:"nome"`
}

//Usuario Handler analisa o request e delega para função adequada
func UsuarioHandler(w http.ResponseWriter, r *http.Request) {
	sid := strings.TrimPrefix(r.URL.Path, "/usuarios/")
	id, _ := strconv.Atoi(sid)

	switch {
	case r.Method == "GET" && id > 0:
		usuarioPorId(w, r, id)
	case r.Method == "GET":
		usuarioTodos(w, r)
	default:
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "Desculpe... :(")
	}
}

func usuarioPorId(w http.ResponseWriter, r *http.Request, id int) {
	db, err := sql.Open("mysql", "root:root@/cursogo")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	var user Usuario
	db.QueryRow("select id, nome from usuarios where id = ?", id).Scan(&user.ID, &user.Nome)

	json, _ := json2.Marshal(user)

	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, string(json))
}

func usuarioTodos(w http.ResponseWriter, r *http.Request) {
	db, err := sql.Open("mysql", "root:root@/cursogo")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	rows, _ := db.Query("select id, nome from usuarios")
	defer rows.Close()

	var usuarios []Usuario

	for rows.Next() {
		var user Usuario
		rows.Scan(&user.ID, &user.Nome)
		usuarios = append(usuarios, user)
	}

	json, _ := json2.Marshal(usuarios)

	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, string(json))
}
