package main

import "fmt"

func main() {

	var employees = map[string]int{"Benjamin": 20, "Nahuel": 26, "Brenda": 19, "Darío": 44, "Pedro": 30}

	fmt.Println(employees["Benjamin"])

	// BUSCO A TODOS LOS MAYORES DE 21
	for empleado, edad := range employees {
		if edad > 21 {
			fmt.Println(empleado, edad)
		}
	}

	// AGREGO A FEDERICO
	employees["Federico"] = 25

	fmt.Println(employees)

	// ELIMINO A PEDRO
	delete(employees, "Pedro")

	fmt.Println(employees)

}
