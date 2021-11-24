/* Ejercicio 1 - Letras de una palabra
La Real Academia Española quiere saber cuántas letras tiene una palabra y luego tener cada una de las letras por separado para deletrearla.
Crear una aplicación que tenga una variable con la palabra e imprimir la cantidad de letras que tiene la misma.
Luego imprimí cada una de las letras.
*/

package main

import (
	"fmt"
	"strings"
)

// Needed to use split

var word string = "Edgar"

func main() {
	fmt.Printf("The word has %v letters. \n", len(word))
	split := strings.Split(word, "")
	fmt.Println(split)
}
