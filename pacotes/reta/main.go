package main

import "fmt"

func main()  {
	p1 := Ponto{2, 2}
	p2 := Ponto{2, 4}

	// visivel apenas no pacote main
	fmt.Println(catetos(p1, p2))
	// visivel para outros pacotes
	fmt.Println(Distancia(p1, p2))
}
