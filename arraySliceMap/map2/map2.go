package main

import "fmt"

func main()  {
	funcionarios := map[string]float64 {
		"José João": 1000.00,
		"Gabriela Silva": 158299.78,
		"Pedro Junior": 1200.0,
	}

	funcionarios["Rafael Filho"] = 1350.0

	delete(funcionarios, "Inexistente")

	for nome, salario := range funcionarios{
		fmt.Println(nome, salario)
	}
}
