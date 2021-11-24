// Ejercicio 4 - A qué mes corresponde

// Realizar una aplicación que contenga una variable con el número del mes.
// 1 Según el número, imprimir el mes que corresponda en texto.
// 2 ¿Se te ocurre si se puede resolver de más de una manera? ¿Cuál elegirías y por qué?
// Ej: 7, Julio
package main

import "fmt"

func main() {

	number_month := 4
	var months = map[int]string{
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

	fmt.Printf("\nMes %d: %v", number_month, months[number_month])
}

// Creo que se puede resolver de más de una manera
// Elijo utilizar un map porque de esta forma obtego un acceso facil al mes mediante su clave.
// Tambien se podría resolver con un switch o con un array
