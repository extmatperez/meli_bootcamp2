package main

import (
	"fmt"
)

func main() {
	f := fibonacci()
	for i := 0; i < 20; i++ {
		fmt.Println(f(i))
	}
}
