package main

import (
	"fmt"
)

func main() {
	var employees = map[string]int{"Benjamin": 20, "Nahuel": 26, "Brenda": 19, "Darío": 44, "Pedro": 30}
	var count int

	fmt.Println("La edad de Benjamín es : ", employees["Benjamin"])

	for _, greater := range employees {
		if greater > 21 {
			count++
		}

	}

	fmt.Println(count, " empleados son mayores de 21 años")

}
