package main

import "fmt"

func main() {
	var employees = map[string]int{"Benjamin": 20, "Nahuel": 26, "Brenda": 19, "Darío": 44, "Pedro": 30}

	fmt.Printf("La edad de Benjamin es %v\n", employees["Benjamin"])

	for employee, age := range employees {
		if age > 21 {
			fmt.Printf("El empleado %v es mayor de 21 años\n", employee)
		}
	}

	employees["Federico"] = 25
	fmt.Printf("Lista con federico: %v\n", employees)

	delete(employees, "Pedro")
	fmt.Printf("Lista sin pedro: %v \n", employees)
}
