package main

import "fmt"

// Ejercicio 6 - Qué edad tiene
// Un empleado de una empresa quiere saber el nombre y edad de uno de sus empleados. Según el siguiente mapa, ayuda  a imprimir la edad de Benjamin. 

// var employees = map[string]int{"Benjamin": 20, "Nahuel": 26, "Brenda": 19, "Darío": 44, "Pedro": 30}

// Por otro lado también es necesario: 
// Saber cuántos de sus empleados son mayores a 21 años.
// Agregar un empleado nuevo a la lista, llamado Federico que tiene 25 años.
// Eliminar a Pedro del mapa.


func main () {

	var employees = map[string]int{"Benjamin": 20, "Nahuel": 26, "Brenda": 19, "Darío": 44, "Pedro": 30}

	name := "Benjamin"
	_, employee_name := employees[name] 

	if employee_name == true {
		fmt.Printf("Nombre: %s\n Edad: %d", name, employees["Benjamin"])

	}


}