package main

import (
	"fmt"
	"math"
	"reflect"
)

func main()  {
	//números inteiros
	fmt.Println(1, 2, 1000)
	fmt.Println("Literal inteiro é", reflect.TypeOf(32000))

	//sem sinal (só positivos)... uint8 uint16 uint32 uint64
	var b byte = 255
	fmt.Println("O byte é", reflect.TypeOf(b))

	//com sinal... int8 int16 int32 int64
	i1 := math.MaxInt64
	fmt.Println("O valor máximo do int é", i1)
	fmt.Println("O byte é", reflect.TypeOf(i1))

	//representação do mapeamento da tabela Unicode (int32)
	var i2 = 'a'
	fmt.Println("O rune é", reflect.TypeOf(i2))
	fmt.Println("O rune é", i2)
	char := 'a'
	fmt.Println("O rune é", reflect.TypeOf(char))
	fmt.Println("O rune é", char)

	//numeros reais float32, float64
	var x float32 = 49.99
	fmt.Println("O tipo de x é", reflect.TypeOf(x))
	fmt.Println("O tipo do literal 49.99 é", reflect.TypeOf(49.99))

	//type boolean
	bo := true
	fmt.Println("O tipo de bo é", reflect.TypeOf(bo))
	fmt.Println(!bo)

	//type string
	s1 := "Olá meu nome é Isaac"
	fmt.Println(s1 + "!")
	fmt.Println("O tamanho da string é", len(s1))

	//string com multiplas linhas
	s2 := `Olá
	meu nome
	é
	Isaac
	`
	fmt.Println(s2)
	fmt.Println("O tamanho da string é", len(s2))
}
