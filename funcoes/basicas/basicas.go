package main

import "fmt"

// função que não recebe nada e não retorna nada
func f1()  {
	fmt.Println("F1")
}

// função com parametro e não retorna nada
func f2(p1 string, p2 string)  {
	fmt.Printf("F2 %s  %s\n", p1, p2)
}

// função que não recebe parametro e retorna string
func f3() string  {
	return "F3"
}

// função resumida que recebe parametro e retorna string
func f4(p1, p2 string) string  {
	return fmt.Sprintf("F4: %s %s", p1, p2)
}

// função que retorna multiplos valores
func f5() (string, string)  {
	return "Retorno 1", "Retorno 2"
}

func main()  {
	f1();
	f2("Teste1", "Teste2")

	r3, r4 := f3(), f4("Teste3", "Teste4")
	fmt.Println(r3, r4);

	r51, r52 := f5()
	fmt.Println(r51, r52);
}