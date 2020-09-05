package main

import "fmt"

func main()  {
	//var aprovados map[int] string
	// Mapas devem ser inicializados!
	aprovados := make(map[int]string)

	aprovados[1234567896] = "Maria"
	aprovados[1232131233] = "Pedro"
	aprovados[2122333122] = "Ana"

	fmt.Println(aprovados)


    for cpf, nome := range aprovados {
    	fmt.Printf("%s (CPF: %d)\n", nome, cpf)
	}

	delete(aprovados, 2122333122)
	fmt.Println(aprovados)
}
