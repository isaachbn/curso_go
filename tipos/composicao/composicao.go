package main

import "fmt"

type esportivo interface {
	ligarTurbo()
}

type luxuoso interface {
	fazerBaliza()
}

type esportivoLuxuoso interface {
	esportivo
	luxuoso
	//pode adicionar mais metodos!
}

type bmw7 struct{}

func (b bmw7) ligarTurbo() {
	fmt.Println("Turbo Ligado!!!")
}
func (b bmw7) fazerBaliza() {
	fmt.Println("Baliza feita!!!")
}

func main() {
	var b esportivoLuxuoso = bmw7{};
	var c luxuoso = bmw7{};
	var d esportivo = bmw7{};

	b.ligarTurbo()
	b.fazerBaliza()

	c.fazerBaliza()

	d.ligarTurbo()
}