package main

import (
	"fmt"
)

var word = "palabra"

func main() {
	fmt.Printf("La cantidad de letras que tiene la palabra es: %d, y sus letras son: \n", len(word))
	for _, r := range word {
		fmt.Printf("%c\n", r)
	}
}
