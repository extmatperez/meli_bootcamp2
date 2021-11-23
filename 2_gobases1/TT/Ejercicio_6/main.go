package main

import "fmt"

func main() {
	var employees = map[string]int{"Benjamin": 20, "Nahuel": 26, "Brenda": 19, "Darío": 44, "Pedro": 30}
	var count int = 0

	for key, value := range employees {
		if key == "Benjamin" {
			fmt.Println("1. La edad de Benjamin es: ", value)

		}
		if value > 21 {
			count++
		}
	}
	fmt.Println("2. Hay ", count, "empleados mayores de 21 años")

	employees["Federico"] = 25
	fmt.Println("3. ", employees)

	delete(employees, "Pedro")
	fmt.Println("4. ", employees)
}
