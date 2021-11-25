package main

import "fmt"

func main() {
	var employees = map[string]int{"Benjamin": 20, "Nahuel": 26, "Brenda": 19, "Darío": 44, "Pedro": 30}

	fmt.Printf("La edad de Benjamin es %d,", employees["Benjamin"])

	var mayores21 int = 0
	for _, edad := range employees {
		if edad > 21 {
			mayores21++
		}
	}
	fmt.Printf("\nHay %d mayores a 21.\n", mayores21)

	employees["Federico"] = 25
	delete(employees, "Pedro")
	fmt.Println(employees)
}
