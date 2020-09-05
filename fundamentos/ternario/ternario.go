package main

import "fmt"

// Não têm operador ternário (nota >=6 ? "Aprovado" : "Reprovado")
func obterResultado(nota float64) string  {
	if nota >=6 {
		return "Aprovado"
	}

	return "Reprovado"
}

func main()  {
	fmt.Println("O aluno foi", obterResultado(6))
}
