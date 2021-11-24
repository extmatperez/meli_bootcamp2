package main

import (
	"fmt"
)

func main() {
	var employees = map[string]int{"Benjamin": 20, "Nahuel": 26, "Brenda": 19, "Darío": 44, "Pedro": 30}
	employees["Federico"] = 25
	delete(employees, "Federico")
	contador := 0
	fmt.Println(employees["Benjamin"])
	for _, element := range employees {

		if element > 21 {
			contador++
		}
	}
	fmt.Println(contador, " empleados mayores de 21")
}
