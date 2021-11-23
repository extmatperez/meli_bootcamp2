package main

import "fmt"

func getEmployeesOver21(employees map[string]int) map[string]int {
	employeesOver21 := make(map[string]int)

	for name, age := range employees {
		if age > 21 {
			employeesOver21[name] = age
		}
	}

	return employeesOver21
}

func main() {
	var employees = map[string]int{"Benjamin": 20, "Nahuel": 26, "Brenda": 19, "Darío": 44, "Pedro": 30}

	fmt.Println("La edad de Benjamin es:", employees["Benjamin"])

	employeesOver21 := getEmployeesOver21(employees)
	fmt.Println("La cantidad de empleados mayor a 21 es:", len(employeesOver21))

	employees["Federico"] = 25
	fmt.Println("Federico agregado a la lista de empleados:", employees)

	delete(employees, "Pedro")
	fmt.Println("Pedro eliminado de la lista de empleados:", employees)
}
