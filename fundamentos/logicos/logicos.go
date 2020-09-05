package main

import "fmt"

func compra(trab1, trab2 bool) (bool, bool, bool)  {
	comprarTv50 := trab1 && trab2
	comprarTv32 := trab1 != trab2 // ou exclusivo
	comprarSorvete := trab1 || trab2

	return comprarTv50, comprarTv32, comprarSorvete
}

func main()  {
	tv50, tv32, sorvete := compra(true, true)
	fmt.Printf("Tv50: %t, Tv32: %t, Sorvete: %t, Saudável: %t\n", tv50, tv32, sorvete, !sorvete)

	tv50, tv32, sorvete = compra(false, true)
	fmt.Printf("Tv50: %t, Tv32: %t, Sorvete: %t, Saudável: %t\n", tv50, tv32, sorvete, !sorvete)

	tv50, tv32, sorvete = compra(false, false)
	fmt.Printf("Tv50: %t, Tv32: %t, Sorvete: %t, Saudável: %t\n", tv50, tv32, sorvete, !sorvete)
}