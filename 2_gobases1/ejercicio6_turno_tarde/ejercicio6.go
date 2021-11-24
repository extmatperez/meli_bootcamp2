// Ejercicio 6 - Qué edad tiene...
// Un empleado de una empresa quiere saber el nombre y edad de uno de sus empleados. Según el siguiente mapa, ayuda  a imprimir la edad de Benjamin.
// Saber cuántos de sus empleados son mayores a 21 años.
// Agregar un empleado nuevo a la lista, llamado Federico que tiene 25 años.
// Eliminar a Pedro del mapa.

//var employees = map[string]int{"Benjamin": 20, "Nahuel": 26, "Brenda": 19, "Darío": 44, "Pedro": 30}

package main

import "fmt"

func main() {
	var employees = map[string]int{"Benjamin": 20, "Nahuel": 26, "Brenda": 19, "Darío": 44, "Pedro": 30}

	fmt.Printf("\nBenjamin, %d \n", employees["Benjamin"])
	count := 0
	for _, age := range employees {
		if age > 21 {
			count++
		}
	}

	fmt.Printf("\nCantidad de empleados mayores a 21 años: %d", count)

	employees["Federico"] = 25

	delete(employees, "Pedro")

	for name, age := range employees {
		println("\n")
		fmt.Println("Nombre:", name, "=>", "Edad:", age)
	}

}
