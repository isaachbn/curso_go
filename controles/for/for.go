package main

import (
	"fmt"
	"time"
)

func main()  {

	//Go não existe while
	i := 1

	for i <= 10 {
		fmt.Print(i);
		i++
	}

	fmt.Println("\n");

	for i:=0; i<=20; i+=2 {
		fmt.Print(i)
	}

	for i:= 1; i<= 10; i++ {
		if i%2 ==0 {
			fmt.Println("Par")
		} else {
			fmt.Println("Impar")
		}
	}

	for {
		//laço infinito
		fmt.Println("Para sempre...")
		time.Sleep(time.Second)
	}
}
