package main

import "fmt"

func main() {

	var employees = map[string]int{"Benjamin": 20, "Nahuel": 26, "Brenda": 19, "Darío": 44, "Pedro": 30}
	cont := 0
	fmt.Println("Empleados:")
	for key, element := range employees {
		fmt.Println("\t", key, "=>", "Edad:", element)
		if element > 21 {
			cont++
		}
	}
	if cont >= 1 {
		fmt.Println("\nNumero de empleados con mas de 21 años:", cont)
	}
	employees["Federico"] = 25
	delete(employees, "Pedro")
	fmt.Println("\nNuevo listado de empleados:")
	for key, element := range employees {
		fmt.Println("\t", key, "=>", "Edad:", element)
		if element > 21 {
			cont++
		}
	}
}
