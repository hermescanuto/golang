package main

import (
	"fmt"
	"hello/sorteio"
)

func main() {
	
	for i:=1 ;i<10 ;i++  {
		fmt.Println("Numero da sorte", sorteio.Randomico(1, 100))
	}
}
