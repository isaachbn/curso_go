package main

import (
	"fmt"
	m "math"
)

func main()  {
	const PI float64 = 3.1415
	// tipo (float64) inferido pelo compilador
	var raio = 3.2
	// forma reduzida de criar um var
	area := PI * m.Pow(raio, 2)

	fmt.Println("A área da circunferência é", area)

	// bloco de construção de const e var
	const (
		a = 1
		b = 2
	)
	var (
		c = 3
		d =4
	)
	fmt.Println(a,b, c, d)

	//Atribuir na mesma linha
	var e, f = true, false
	g, h, i := 2, false, "ops!"
	fmt.Println(e, f, g, h, i)
}
