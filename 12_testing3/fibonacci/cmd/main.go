package main

import (
	"fmt"

	L "github.com/extmatperez/meli_bootcamp2/12_testing3/fibonacci/internal/fibonacci"
)

func main() {
	f := L.Fibonacci()
	for i := 0; i < 20; i++ {
		fmt.Println(f(i))
	}
}
