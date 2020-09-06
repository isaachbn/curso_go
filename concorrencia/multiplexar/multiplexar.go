package main

import (
	"fmt"
	"github.com/isaachbn/html"
)

func encaminhar(origem <-chan string, destino chan string) {
	for {
		destino <- <-origem
	}
}

// multiplexar - misturar(mensagens) num canal
func juntar(entrada1, entrada2 <-chan string) <-chan string {
	c := make(chan string)
	go encaminhar(entrada1, c)
	go encaminhar(entrada2, c)

	return c
}

func main() {
	c := juntar(
		html.Titulo("https://www.google.com", "https://www.youtube.com"),
		html.Titulo("https://www.github.com", "https://www.cod3r.com.br"),
	)

	fmt.Println(<-c, "\n", <-c)
	fmt.Println(<-c, "\n", <-c)
}
