package main

import "fmt"

// Ejercicio 1 - Letras de una palabra

func main () {

	word := "golang"
	word_length := len(word)

	fmt.Printf("La cantidad de letras es: %d", word_length)

	for position, letter := range word {
		fmt.Printf("Letra %d: %s\n", position, string(letter))
	}

}