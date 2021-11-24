package main

import "fmt"

func main() {

	var employees = map[string]int{"Benjamin": 20, "Nahuel": 26, "Brenda": 19, "Darío": 44, "Pedro": 30}

	fmt.Printf("Edad de Benjamin: %d\n", employees["Benjamin"])

	employeesOlderThan21Years := 0

	for _, age := range employees {
		if age > 21 {
			employeesOlderThan21Years++
		}
	}

	fmt.Printf("Empleados mayores a 21: %d\n", employeesOlderThan21Years)

	fmt.Printf("%v\n", employees)
	employees["Federico"] = 25
	delete(employees, "Pedro")
	fmt.Printf("%v\n", employees)

}
