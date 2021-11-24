package main

import "fmt"

func main() {

	var employees = map[string]int{"Benjamin": 20, "Nahuel": 26, "Brenda": 19, "Darío": 44, "Pedro": 30}

	name := "Nahuel"

	fmt.Println("La edad de ", name, " es ", employees[name])

	var greater_than_21 int = 0
	for _, valor := range employees {
		if valor > 21 {
			greater_than_21++
		}
	}
	fmt.Printf("Hay %v empleados mayores de 21 años \n", greater_than_21)

	fmt.Println("Empleados: ", employees)

	employees["Federico"] = 25

	fmt.Println("Empleados: ", employees)

	delete(employees, "Pedro")

	fmt.Println("Empleados: ", employees)

}
