package main

import "fmt"

func main()  {
	i := 1

	//Go não têm aritmética de ponteiros
	var p *int = nil

	p = &i //atribuindo o endereço de memoria de i

	*p++ // desreferenciando - incrementando o valor da variavel de i
	i++

	fmt.Println(p, *p, i, &i)
}
