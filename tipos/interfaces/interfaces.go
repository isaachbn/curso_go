package main

import "fmt"

type imprimivel interface {
	toString() string
}

type pessoa struct {
	nome      string
	sobrenome string
}

type produto struct {
	nome  string
	preco float64
}

//interfaces são implementadas implicitamente
func (p pessoa) toString() string {
	return p.nome + " " + p.sobrenome
}

func (p produto) toString() string {
	return fmt.Sprintf("%s - R$ %.2f", p.nome, p.preco)
}

func imprimir(x imprimivel)  {
	fmt.Println(x.toString())
}

func main() {
	var coisa imprimivel = pessoa{nome: "Lorena", sobrenome: "Farias"}
	imprimir(coisa)
	coisa = produto{nome: "Caixa de Som", preco: 48.76}
	imprimir(coisa)

	imprimir(pessoa{nome: "Isaac", sobrenome: "Henrique"})
	imprimir(produto{nome: "TCL QLEd 55", preco: 3780.80})
}