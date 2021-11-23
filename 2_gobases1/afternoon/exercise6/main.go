package main

import "fmt"

func main() {
	/*
		Un empleado de una empresa quiere saber el nombre y edad de uno de sus empleados. Según el siguiente mapa, ayuda  a imprimir la edad de Benjamin.

		var employees = map[string]int{"Benjamin": 20, "Nahuel": 26, "Brenda": 19, "Darío": 44, "Pedro": 30}

		Por otro lado también es necesario:

		- Saber cuántos de sus empleados son mayores a 21 años.
		- Agregar un empleado nuevo a la lista, llamado Federico que tiene 25 años.
		- Eliminar a Pedro del mapa.
	*/

	var employees = map[string]int{"Benjamin": 20, "Nahuel": 26, "Brenda": 19, "Darío": 44, "Pedro": 30}
	counter := 0

	fmt.Println(employees["Benjamin"])

	for _, value := range employees {
		if value > 21 {
			counter++
		}
	}
	fmt.Println("La cantidad de empleados mayores a 21 años es: ", counter)

	employees["Federico"] = 25

	fmt.Println(employees)

	delete(employees, "Pedro")

	fmt.Println(employees)
}
