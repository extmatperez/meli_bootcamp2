package main

import "fmt"

func main() {

	var palabra = "palabra de prueba"

	fmt.Println("la canitdad de letras de '", palabra, "' es de: ", len(palabra))

	for i := 0; i < len(palabra); i++ {
		fmt.Printf("\n%c", palabra[i])
	}
}
