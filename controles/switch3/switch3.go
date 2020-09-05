package main

import (
	"fmt"
	"time"
)

func tipo(i interface{}) string {
	switch i.(type) {
	case int:
		return "inteiro"
	case string:
		return "string"
	case float32, float64:
		return "real"
	case bool:
		return "boolean"
	case func():
		return "função"
	default:
		return "não sei"
	}
}

func main()  {
	fmt.Println(tipo(2.3))
	fmt.Println(tipo(2))
	fmt.Println(tipo("isaac"))
	fmt.Println(tipo(func() {}))
	fmt.Println(tipo(true))
	fmt.Println(tipo(time.Now()))
}
