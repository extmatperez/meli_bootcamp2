package main

import "fmt"

func main() {
	/*
		Realizar una aplicación que contenga una variable con el número del mes.
		1 - Según el número, imprimir el mes que corresponda en texto.
		2 - ¿Se te ocurre si se puede resolver de más de una manera? ¿Cuál elegirías y por qué?
	*/

	number := 9

	// Opcion 1

	switch number {
	case 1:
		fmt.Println("Enero")
	case 2:
		fmt.Println("Febrero")
	case 3:
		fmt.Println("Marzo")
	case 4:
		fmt.Println("Abril")
	case 5:
		fmt.Println("Mayo")
	case 6:
		fmt.Println("Junio")
	case 7:
		fmt.Println("Julio")
	case 8:
		fmt.Println("Agosto")
	case 9:
		fmt.Println("Septiembre")
	case 10:
		fmt.Println("Octubre")
	case 11:
		fmt.Println("Noviembre")
	case 12:
		fmt.Println("Diciembre")
	default:
		fmt.Println("No es un mes valido")
	}

	// Opcion 2

	months2 := [12]string{"Enero", "Febrero", "Marzo", "Abril", "Mayo", "Junio", "Julio", "Agosto", "Septiembre", "Octubre", "Noviembre", "Diciembre"}

	fmt.Println(months2[number-1])

	// Opcion 3

	months3 := map[int]string{
		1:  "Enero",
		2:  "Febrero",
		3:  "Marzo",
		4:  "Abril",
		5:  "Mayo",
		6:  "Junio",
		7:  "Julio",
		8:  "Agosto",
		9:  "Septiembre",
		10: "Octubre",
		11: "Noviembre",
		12: "Diciembre",
	}

	fmt.Println(months3[number])

	/*
		De las opciones encontradas, la que mas me gusta es la opcion 2, porque es la mas sencilla y facil de entender.
	*/
}
