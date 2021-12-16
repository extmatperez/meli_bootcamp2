package main

import "fmt"

func main() {
	f := internal.fibonacci()
	for i := 0; i < 20; i++ {
		fmt.Println(f(i))
	}
}
