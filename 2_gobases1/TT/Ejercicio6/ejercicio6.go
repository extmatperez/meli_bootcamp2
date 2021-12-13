package main

import "fmt"

func main() {
	var employees = map[string]int{"Benjamin": 20, "Nahuel": 26, "Brenda": 19, "Darío": 44, "Pedro": 30}
	fmt.Println(employees["Benjamin"])
	older_than_twentyone(employees)
	employees["Federico"] = 25
	delete(employees, "Pedro")
	older_than_twentyone(employees)
}

func older_than_twentyone(empleados map[string]int) {
	for k, e := range empleados {
		if e > 21 {
			fmt.Println(k)
		}
	}
}
