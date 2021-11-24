package main

import "fmt"

func main() {
	var employees = map[string]int{"Benjamin": 20, "Nahuel": 26, "Brenda": 19, "Darío": 44, "Pedro": 30}
	var count int = 0

	fmt.Println("Edad de Benjamin:",employees["Benjamin"])

	for _, value := range employees {
		if value > 21{
			count += 1
		}
	}

	fmt.Println("Hay",count,"empleados mayores de 21 años")

	fmt.Println(employees)

	employees["Federico"] = 21

	fmt.Println(employees)

	delete(employees, "Pedro")

	fmt.Println(employees)
}