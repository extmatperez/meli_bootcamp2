package main

import "fmt"

func main() {
	var employees = map[string]int{"Benjamin": 20, "Nahuel": 26, "Brenda": 19, "Dario": 44, "Pedro": 30}
	cont := 0
	for key, element := range employees {
		if element > 21 {
			cont++
		}
		_ = key
	}
	fmt.Println("Cantidad de alumnos con edad mayor a 21 : ", cont)
	employees["Federico"] = 25
	delete(employees, "Pedro")
	fmt.Println(employees)

}
