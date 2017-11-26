package sorteio

import (
	"math/rand"
	"time"
)

func Randomico(min, max int) int {
	rand.Seed(time.Now().UTC().UnixNano())
	return rand.Intn(max-min) + min
}
