package main

import "fmt"

func main()  {
	fmt.Print("Mesma")
	fmt.Print(" linha\n")

	fmt.Println("Nova")
	fmt.Println("lina.")

	x := 3.141516
	xs := fmt.Sprint(x)
	fmt.Println("O valor de x é "+ xs)
	fmt.Println("O valor de x é", x)

	/**
	  %f imprime do tipo float
	 */
	fmt.Printf("O valor de x é %f\n", x)
	fmt.Printf("O valor de x é %.2f", x)

	a, b, c, d := 1, 1.9999, false, "ops!"

	/**
	  %d imprimir inteiro
	  %f imprimir float
	  %t imprimir boolean
	  %s imprimir string
	  %q imprimir default
	  %v generic
	 */
	fmt.Printf("\n %d %f %.2f %t %s", a, b, b, c, d)
	fmt.Printf("\n %v %v %.2v %v %v", a, b, b, c, d)
}
