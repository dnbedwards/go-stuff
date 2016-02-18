package main

import (
	"fmt"
	"math/rand"
	"time"
)

func randomInts(max int) (int, int) {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(max), rand.Intn(max)
}

func main() {
	a, b := randomInts(10)
	fmt.Println("My favourite numbers are", a, b)
}
