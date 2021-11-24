package main

import "fmt"

var employees = map[string]int{"Benjamin": 20, "Nahuel": 26, "Brenda": 19, "Dario": 44, "Pedro": 30}

func main() {
	counter := 0

	fmt.Printf("\n %v", employees["Benjamin"])
	for _, value := range employees {
		if value > 21 {
			counter++
		}
	}
	fmt.Printf("\nLa cantidad de empleados mayores a 21 son: %v", counter)

	employees["Federico"] = 25
	delete(employees, "Pedro")
}
