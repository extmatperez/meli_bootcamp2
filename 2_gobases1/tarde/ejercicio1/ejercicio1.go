package main

import "fmt"

func letrasPalabra(palabra string) {
	var cantLetras = len(palabra)
	var deletreo []string
	for i := 0; i < cantLetras; i++ {
		deletreo = append(deletreo, string([]rune(palabra)[i]))
	}
	fmt.Printf("La cantidad de letras que tiene la palabra %v es: %v\n", palabra, cantLetras)
	fmt.Printf("La palabra %v deletreada es: %v", palabra, deletreo)
}

func main() {
	letrasPalabra("ivan")
}
