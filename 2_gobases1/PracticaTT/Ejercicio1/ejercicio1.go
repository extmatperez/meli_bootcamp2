package main

import "fmt"

func main() {
	var palabra = "terremoto"
	var largoPalabra = len(palabra)
	fmt.Println("El largo de la palabra es", largoPalabra)

	for i := 0; i < largoPalabra; i++ {
		fmt.Printf("%c\n", palabra[i])
	}
}
