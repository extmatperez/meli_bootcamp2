package main

import (
	"fmt"
	"strings"
)

/*
La Real Academia Española quiere saber cuántas letras tiene una palabra
y luego tener cada una de las letras por separado para deletrearla.
	1 - Crear una aplicación que tenga una variable con la palabra e imprimir
 la cantidad de letras que tiene la misma.
	2 - Luego imprimí cada una de las letras.
*/

func main() {
	palabra := "palabra"

	cadena := strings.Split(palabra, "")

	x := 0
	for i, letra := range cadena {
		fmt.Println(i, letra)
		x = i
	}

	for _, letter := range palabra {
		//fmt.Println(string(letter))
		fmt.Printf("\n%c", letter)
	}

	fmt.Printf("La palabra tiene: %d letras\n", x)
}
