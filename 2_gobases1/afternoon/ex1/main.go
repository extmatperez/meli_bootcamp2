package main

import "fmt"

// Letras de una palabra
func main() {

	word := "MercadoPago"

	fmt.Println("Cantidad de letras:", len(word))

	for i := 0; i < len(word); i++ {
		fmt.Printf("%c\n", word[i]);
	}

}