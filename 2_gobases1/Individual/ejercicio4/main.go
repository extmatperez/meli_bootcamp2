/*Realizar una aplicación que contenga una variable con el número del mes.
Según el número, imprimir el mes que corresponda en texto.
¿Se te ocurre si se puede resolver de más de una manera? ¿Cuál elegirías y por qué?
*/
package main

import (
	"fmt"
)

func main() {

	var mes int = 5

	switch mes {
	case 1:
		fmt.Printf("El mes es enero\n")
	case 2:
		fmt.Printf("El mes es febrero\n")
	case 3:
		fmt.Printf("El mes es marzo\n")
	case 4:
		fmt.Printf("El mes es abril\n")
	case 5:
		fmt.Printf("El mes es mayo\n")
	case 6:
		fmt.Printf("El mes es junio\n")
	case 7:
		fmt.Printf("El mes es julio\n")
	case 8:
		fmt.Printf("El mes es agosto\n")
	case 9:
		fmt.Printf("El mes es septiembre\n")
	case 10:
		fmt.Printf("El mes es octubre\n")
	case 11:
		fmt.Printf("El mes es noviembre\n")
	case 12:
		fmt.Printf("El mes es diciembre\n")
	default:
		fmt.Printf("El mes ingresado es incorrecto")
	}

}
