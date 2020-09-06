package main

import (
	"fmt"
	"time"
)

// Channel (canal) - é a forma de comunicação entre goroutines
// Channel é um tipo

func doisTresQuatroVezes(base int, c chan int) {
	time.Sleep(time.Second)
	c <- 2 * base // enviando dados para o canal, só para para proxima interação quando esse valor for lido.
	time.Sleep(time.Second)
	c <- 3 * base
	time.Sleep(3 * time.Second)
	c <- 4 * base
}
func main() {
	// Criando channel sem buffer
	ch := make(chan int)
	go doisTresQuatroVezes(2, ch)

	a, b := <-ch, <-ch // recebendo os dados do canal com sincronismo
	fmt.Println(a, b)

	fmt.Println(<-ch)
}
