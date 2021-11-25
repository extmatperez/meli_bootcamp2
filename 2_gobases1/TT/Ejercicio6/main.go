package main

import "fmt"

func main() {
	var employees = map[string]int{"Benjamin": 20, "Nahuel": 26, "Brenda": 19, "Dario": 44, "Pedro": 30}

	counter := 0

	fmt.Println(employees["Benjamin"])

	for _, value := range employees {
		if value > 21 {
			counter++
		}
	}

	fmt.Println("La cantidad de empleados mayores a 21 es ", counter)

	employees["Federico"] = 25

	fmt.Println(employees)

	delete(employees, "Pedro")

	fmt.Println(employees)
}
