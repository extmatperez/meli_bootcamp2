package main

import "fmt"

func main() {

	var palabra string
	palabra = "ejercicio"

	fmt.Println("Cantidad de letras: ", len(palabra))
	fmt.Println("letras: \n")
	for i := 0; i < len(palabra); i++ {
		fmt.Printf("%c \n", palabra[i])
	}

}
