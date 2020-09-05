package main

import (
	"fmt"
	"math"
)

func main()  {
	a := 3
	b := 2

	fmt.Println("Soma =>", a+b)
	fmt.Println("Subtração =>", a-b)
	fmt.Println("Divisão =>", a/b)
	fmt.Println("Multiplicação =>", a*b)
	fmt.Println("Módulo =>", a%b)

	// bitwise
	fmt.Println("E =>", a&b) // 11 & 10 = 10
	fmt.Println("OU =>", a|b) // 11 | 10 = 11
	fmt.Println("Xor =>", a^b) // 11 ^ 10 = 01

	// Outras operações usando o math...
	c := 3.0
	d := 2.0
	//maior valor entre dois numeros
	fmt.Println("Maior =>", math.Max(float64(a), float64(b)))
	fmt.Println("Maior =>", math.Max(c, d))
	//menor valor entre dois numeros
	fmt.Println("Menor =>", math.Min(c, d))
	//exponencial
	fmt.Println("Exponenciação =>", math.Pow(c, d))
}
