package main

import "fmt"

func fatorial(n int) (int, error) {
	switch {
	case n < 0:
		return -1, fmt.Errorf("número inválidos: %d", n)
	case n == 0:
		return 1, nil
	default:
		fatorialAnterior, _ := fatorial(n - 1)
		return n * fatorialAnterior, nil
	}
}

func fatorialSimple(n uint) uint {
	switch {
	case n == 0:
		return 1
	default:
		return n * fatorialSimple(n - 1)
	}
}

func main()  {
	fmt.Println(fatorial(2))
	fmt.Println(fatorial(-1))
	fmt.Println(fatorialSimple(5))
}
