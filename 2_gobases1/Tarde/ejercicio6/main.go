/*
Un empleado de una empresa quiere saber el nombre y edad de uno de sus empleados.
Según el siguiente mapa, ayuda  a imprimir la edad de Benjamin.

var employees = map[string]int{"Benjamin": 20, "Nahuel": 26, "Brenda": 19, "Darío": 44, "Pedro": 30}

Por otro lado también es necesario:
Saber cuántos de sus empleados son mayores a 21 años.
Agregar un empleado nuevo a la lista, llamado Federico que tiene 25 años.
*/

package main

import "fmt"

func main() {

	var employees = map[string]int{"Benjamin": 20, "Nahuel": 26, "Brenda": 19, "Darío": 44, "Pedro": 30}

	employees["Federico"] = 25

	contador :=  0

	for _, edad := range employees {
		if edad > 21 {
			contador++
		}
	}

	fmt.Println(employees)
	fmt.Println("La edad de Benjamin es", employees["Benjamin"])
	fmt.Println("La cantidad de empleados mayores a 21 años es:", contador)
}
