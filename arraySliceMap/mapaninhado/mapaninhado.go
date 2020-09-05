package main

import "fmt"

func main() {
	funcsPorLetra := map[string]map[string]float64{
		"Marido": {
			"Isaac": 4900.00,
		},
		"Esposa": {
			"Lorena": 15000.80,
		},
		"Animal": {
			"Cachorro": -122,
		},
	}

	delete(funcsPorLetra, "Animal")

	for tipo, pessoa := range funcsPorLetra {
		fmt.Println(tipo, pessoa)
		for _, valor := range pessoa {
			fmt.Println(valor)
		}
	}
}
