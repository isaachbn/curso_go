package main

import "fmt"

func main() {
	// criando um canal de inteiro com buffer de 1
	ch := make(chan int, 1)

	ch <- 1 // enviando dados para o canal(escrita)
	<-ch// recebendo dados do canal(leitura)

	ch <- 2
	fmt.Println(<-ch)
}
