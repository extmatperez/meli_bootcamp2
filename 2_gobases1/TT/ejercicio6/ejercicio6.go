package main

import "fmt"

func main() {
	var employees = map[string]int{"Benjamin": 20, "Nahuel": 26, "Brenda": 19, "Darío": 44, "Pedro": 30}
	fmt.Printf("Benjamin tiene %v años\n",employees["Benjamin"])

	over21 := 0
	for _, employee := range employees {
		if employee >= 21 {
			over21++
		}
	}
	fmt.Printf("Hay %v empleados mayores de 21 años\n", over21)

	employees["Federico"] = 25
	delete(employees, "Pedro")
	fmt.Println(employees)
}