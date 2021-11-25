package main

import "fmt"

func main() {
	var palabra string // = "hola"

	fmt.Printf("Ingrese la palabra: ")
	fmt.Scanf("%s", &palabra)

	var cantidad int = len(palabra)

	fmt.Printf("cantidad: %d\n", cantidad)

	for i := 0; i < cantidad; i++ {
		fmt.Printf("%c\n", palabra[i]) // %c para char
	}

	//otra forma:

	// for _, letra := range palabra {
	// 	println(string(letra))
	// }
}
