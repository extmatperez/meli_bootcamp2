package main

import "fmt"

func main() {

	var employees = map[string]int{"Benjamin": 20, "Nahuel": 26, "Brenda": 19, "Darío": 44, "Pedro": 30}
	fmt.Println("Edad Benjamin:", employees["Benjamin"])
	fmt.Println("Mayores de 21:")

	for nombre, edad := range employees {
		if edad > 21 {
			fmt.Println("Nombre:", nombre, "edad:", edad)
		}
	}
	employees["Federico"] = 25
	delete(employees, "Pedro")
	fmt.Println("Eliminar a Pedro y agregar a Federico:")
	fmt.Println(employees)
}
