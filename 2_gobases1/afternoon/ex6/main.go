package main

import "fmt"

// Qué edad tiene...
func main() {

	var employees = map[string]int{"Benjamin": 20, "Nahuel": 26, "Brenda": 19, "Darío": 44, "Pedro": 30}

	fmt.Println("Edad de Benjamin:", employees["Benjamin"])

	counter := 0
	for _, v := range employees {
		if v > 21 {
			counter++
		}
	}

	fmt.Println("La cantidad de empleados mayores a 21 es:", counter)

	employees["Federico"] = 25

	delete(employees, "Pedro")

	fmt.Println(employees)

}
