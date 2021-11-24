package main

import "fmt"

func main() {

	palabra := "bases"

	fmt.Printf("cantidad de letras: %d\n", len(palabra))

	for _, char := range palabra {
		fmt.Printf("%c\n", char)
	}
}
