package main

import (
	"fmt"
)

func main() {
	var employees = map[string]int{"Benjamin": 20, "Nahuel": 26, "Brenda": 19, "Darío": 44, "Pedro": 30}
	fmt.Printf("Edad Benjamin:%v", employees["Benjamin"])
	fmt.Println("\nMayores de 21:")
	for nombre, edad := range employees {
		if edad > 21 {
			fmt.Printf("Nombre:%v edad:%v\n", nombre, edad)
		}
	}
	employees["Federico"] = 25
	delete(employees, "Pedro")
	fmt.Println("\nEliminar a Pêdro y agregar a Federico:")
	fmt.Println(employees)
}
