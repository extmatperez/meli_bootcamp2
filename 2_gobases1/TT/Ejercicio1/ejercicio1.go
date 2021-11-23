package main

import "fmt"

func main() {
	palabra := "Palabra"

	fmt.Printf("La cantidad de caracteres de la palabra es: %d\n", len(palabra))

	for i := 0; i < len(palabra); i++ {
		fmt.Printf("%c\n", palabra[i])
	}
}
