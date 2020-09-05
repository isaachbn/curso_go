package main

import "fmt"

type item struct {
	produtoId  int
	quantidade int
	preco      float64
}

type pedido struct {
	userId int
	itens  []item
}

func (p pedido) valorTotal() float64 {
	total := 0.0
	for _, item := range p.itens {
		total += item.preco * float64(item.quantidade)
	}

	return total
}

func main() {
	pedido := pedido{
		userId: 1,
		itens: []item{
			item{1, 2, 12.10},
			{2, 1, 23.49},
			item{preco: 100.50, quantidade: 4, produtoId: 3},
		},
	}

	fmt.Printf("Valor totoal do pedido é R$ %.2f", pedido.valorTotal())
}
