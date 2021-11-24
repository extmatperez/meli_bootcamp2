// Ejercicio 1 - Letras de una palabra
// La Real Academia Española quiere saber cuántas letras tiene una palabra y luego tener cada una de las letras por separado para deletrearla.
// 1 Crear una aplicación que tenga una variable con la palabra e imprimir la cantidad de letras que tiene la misma.
// 2 Luego imprimí cada una de las letras.

package main

import "fmt"

func main() {
	word := "Hola"
	println("\n")
	fmt.Println("La cantidad de letras que tiene la palabra es: ", len(word))
	fmt.Println("\nSus letras son: ")

	for _, char := range word {
		fmt.Printf("\t\t%c \n", char)
	}
}
