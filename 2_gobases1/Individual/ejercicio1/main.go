/*La Real Academia Española quiere saber cuántas letras tiene una palabra y luego tener cada una de las letras por separado
para deletrearla.
Crear una aplicación que tenga una variable con la palabra e imprimir la cantidad de letras que tiene la misma.
Luego imprimí cada una de las letras.
*/

package main

import (
	"fmt"
)

func main() {
	var word string = "golang"
	var long int
	long = len(word)

	fmt.Printf("La palabra %s tiene %d letras\n", word, long)

	for _, letter := range word {
		fmt.Printf("%c  ", letter)
	}
	fmt.Printf("\n")

}
