package main

import (
	"fmt"
	"time"
)

func main()  {
	//strings
	fmt.Println("String ==", "Banana" == "Banana")

	//numeros
	fmt.Println("!=", 3 != 2)
	fmt.Println("<", 3 < 2)
	fmt.Println(">", 3 > 2)
	fmt.Println("<=", 3 <= 2)
	fmt.Println(">=", 3 >= 2)

	//datas
	d1 := time.Unix(0, 0)
	d2 := time.Unix(0, 0)
	fmt.Println("Datas ==", d1 == d2)
	fmt.Println("Datas equal", d1.Equal(d2))

	//struct
	type Pessoa struct {
		Nome string
	}
	p1 := Pessoa{"Isaac"}
	p2 := Pessoa{"Isaac"}
	fmt.Println("Struct ==", p1 == p2)
}
