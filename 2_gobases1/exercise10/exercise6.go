package main

import "fmt"

func main() {
	var count int
	var employees = map[string]int{"Benjamin": 20, "Nahuel": 26, "Brenda": 19, "Darío": 44, "Pedro": 30}
	fmt.Println("La edad de Benjamin es: ", employees["Benjamin"])

	for key, element := range employees {
		if element >= 21 {
			fmt.Println("EL empleado ", key, " tiene sobre 21 años")
			count++
		}
	}
	fmt.Println("La cantidad de empleados mayores a 21 años son: ", count)

	fmt.Println("Actual canidad de empleados: ", len(employees), " se agregara a Federico.")
	employees["Federico"] = 25
	fmt.Println("Actual canidad de empleados: ", len(employees), " se eliminara a Pedro.")
	delete(employees, "Pedro")
	fmt.Println("Actual canidad de empleados: ", len(employees))
}
