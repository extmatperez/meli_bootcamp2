/*Un empleado de una empresa quiere saber el nombre y edad de uno de sus empleados.
Según el siguiente mapa, ayuda  a imprimir la edad de Benjamin.

  var employees = map[string]int{"Benjamin": 20, "Nahuel": 26, "Brenda": 19, "Darío": 44, "Pedro": 30}

Por otro lado también es necesario:
Saber cuántos de sus empleados son mayores a 21 años.
Agregar un empleado nuevo a la lista, llamado Federico que tiene 25 años.
Eliminar a Pedro del mapa.*/

package main

import "fmt"

func main() {

	var employees = map[string]int{"Benjamin": 20, "Nahuel": 26, "Brenda": 19, "Darío": 44, "Pedro": 30}
	var cant_mayores int = 0

	if employees["Benjamin"] != 0 {
		fmt.Println("La edad de Benjamin es: ", employees["Benjamin"])
	}

	for _, elem := range employees {
		if elem > 21 {
			cant_mayores++
		}
	}

	fmt.Println("Cantiadad de Empleado mayores de 21: ", cant_mayores)
	fmt.Println("")
	fmt.Println("Lista Empleados: ")

	for i, elem := range employees {
		fmt.Println(i, elem)
	}

	employees["Federico"] = 25

	delete(employees, "Pedro")
	fmt.Println("")
	fmt.Println("Lista Empleados Modificada: ")

	for i, elem := range employees {
		fmt.Println(i, elem)
	}
}
