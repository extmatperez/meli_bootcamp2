package main

import "fmt"

func main() {
	word := "Prueba"
	length := len(word)

	fmt.Println("La palabra ", word," tiene ",length , " letras de largo.")
	
	for i := 0; i < len(word); i++ {
		fmt.Println("La letra en la posicion", i , " es ", string(word[i]))
	}
}