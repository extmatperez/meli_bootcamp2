package main

import "fmt"

func main() {

	var employees = map[string]int{"Benjamin": 20, "Nahuel": 26, "Brenda": 19, "Darío": 44, "Pedro": 30}

	fmt.Println("La edad de Benjamin es: ", employees["Benjamin"])
	count := 0
	for _, i := range employees {
		if i > 21 {
			count = count+1
		}
	}
	fmt.Println("La cantidad de empleados mayores a 21 es: ", count)

	employees["Federico"] = 25
	delete(employees, "pedro")

	fmt.Println("La nueva lista de empleados es: ", employees)
}