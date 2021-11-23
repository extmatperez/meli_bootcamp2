package main

import "fmt"

var employees = map[string]int{"Benjamin": 20, "Nahuel": 26,
	"Brenda": 19, "Dario": 44, "Pedro": 30,
}

func count() int {
	quantity := 0
	for _, age := range employees {
		if age >= 21 {
			quantity = quantity + 1
		}
	}
	return quantity
}

func main() {
	fmt.Println("Amount of employees who are older than 21: ", count())
	employees["Federico"] = 25
	delete(employees, "Pedro")
}
