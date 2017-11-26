package main

import (
	"fmt"
	"math/rand"
	"time"
)

func sorteio(min, max int) int {
	rand.Seed(time.Now().UTC().UnixNano())
	return rand.Intn(max-min) + min
}

func main() {

	fmt.Println("Numero da sorte", sorteio(1, 100))

}
