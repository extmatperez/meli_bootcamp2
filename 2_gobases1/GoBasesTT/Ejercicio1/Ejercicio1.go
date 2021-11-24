package main

import "fmt"

func main() {
	letrasDeUnaPlabra()
}

func letrasDeUnaPlabra() {
	var palabra string = "Hola"

	fmt.Printf("Cantidad de letras: %d \n", len(palabra))

	for i := 0; i < len(palabra); i++ {
		fmt.Printf("\n Letra :%c", (palabra[i]))
	}
}
