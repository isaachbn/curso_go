package main

import "fmt"

func rotina(ch chan int)  {
	ch <- 1
	ch <- 2
	ch <- 3
	ch <- 4
	ch <- 5
	fmt.Println("Não executa mais...")
	ch <- 6
	ch <- 7
}

func main() {
	ch := make(chan int, 3) // canal com 3 possições de buffer

	go rotina(ch)

	fmt.Println(<-ch)
}
