package main

import "fmt"

func main() {
	var employees = map[string]int{"Benjamin": 20, "Nahuel": 26, "Brenda": 19, "Darío": 44, "Pedro": 30}
	fmt.Printf("Edad de Benjamin: %d\n", employees["Benjamin"])

	var aux int

	// Ahora se va a imprimir la cantidad de los empleados mayores a 21 años.
	for key := range employees {
		if employees[key] >= 21 {
			aux++
		}
	}
	fmt.Printf("Cantidad de empleados mayores a 21 años: %d\n", aux)
	employees["Federico"] = 25
	delete(employees, "Pedro")
	fmt.Println("Listado de empleados último: ")
	// Listado sin Pedro y con Federico.
	for key, element := range employees {
		fmt.Printf("%s %d\n", key, element)
	}
}
