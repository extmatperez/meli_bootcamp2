package main

import "fmt"

func main() {
	var employees = map[string]int{"Benjamin": 20, "Nahuel": 26, "Brenda": 19, "Darío": 44, "Pedro": 30}
	fmt.Printf("La edad de Benjamin es %d\n", employees["Benjamin"])
	var mayores21 = 0
	for _, value := range employees {
		if value >= 21 {
			mayores21++
		}
	}
	fmt.Printf("Empleados mayores de 21 años: %d", mayores21);
	employees["Federico"] = 25
	delete(employees, "Pedro")
	fmt.Printf("\nEmpleados: %v", employees)
}