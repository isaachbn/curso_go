package main

import (
	"fmt"
	"github.com/isaachbn/html"
)

func main() {
	t1 := html.Titulo("https://www.google.com", "https://www.youtube.com")
	t2 := html.Titulo("https://www.github.com", "https://www.cod3r.com.br")

	fmt.Println("Primeiros: ", <-t1, "|", <-t2)
	fmt.Println("Segundos: ", <-t1, "|", <-t2)
}
