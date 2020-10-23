package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func horaCerta(w http.ResponseWriter, r *http.Request)  {
	/**
	  02 - representa dia
	  01 - representa mês
	  2006 - reprrsenta ano
	  03 - representa hora
	  04 - representa minuto
	  05 - representa segundo
	**/
	s := time.Now().Format("02/01/2006 03:04:05")
	fmt.Fprintf(w, "<h1>HORA CERT: %s</h1>", s)
}

func main() {
	http.HandleFunc("/horaCerta", horaCerta)
	log.Println("Executando")
	log.Fatal(http.ListenAndServe(":3000", nil))
}