package main

import (
	"fmt"
	"strings"
)

func main() {

	word := "manzana"
	fmt.Println("La palabra tiene", len(word), "letras")

	palabra := strings.Split(word, "")

	fmt.Print(palabra)
	fmt.Println("\n")
	for i := 0; i < len(word); i++ {

		fmt.Println(palabra[i])

	}
}
