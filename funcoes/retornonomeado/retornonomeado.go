package main

import "fmt"

// retorno nomeado
func troca(p1, p2 int) (segundo, primeiro int)  {
	segundo = p2
	primeiro = p1
	return //retorno limpo
}

func main()  {
	fmt.Println(troca(1, 2))
}
