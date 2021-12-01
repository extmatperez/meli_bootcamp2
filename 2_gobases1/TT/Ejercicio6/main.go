package main

import "fmt"

func main() {
	var employees = map[string]int{"Benjamin": 20, "Nahuel": 26, "Brenda": 19, "Darío": 44, "Pedro": 30}
	employees_greater_21 := 0
	fmt.Println("La edad de Benjamin es:",employees["Benjamin"], "años")

	for key, element := range employees {
		if element > 21{
			employees_greater_21 += 1
			fmt.Printf("La edad de %v es : %v años\n",key,element)
		}

	}
	fmt.Println("La cantidad de empleados mayores a 21 es:", employees_greater_21)

	employees["Federico"] = 25

	delete(employees, "Pedro")
	fmt.Println(employees)
}