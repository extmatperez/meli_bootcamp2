package main

import "fmt"

func main() {

	var employees = map[string]int{"Benjamin": 20, "Nahuel": 26, "Brenda": 19, "Darío": 44, "Pedro": 30}

	fmt.Println("Benjamin's age is:", employees["Benjamin"])

	var biggerOfTwentyOne = map[string]int{}
	for key, value := range employees {
		if value > 21 {
			biggerOfTwentyOne[key] = value
		}
	}

	fmt.Printf("%d employees over 21 years of age\n", len(biggerOfTwentyOne))

	fmt.Println(employees)
	employees["Federico"] = 25
	delete(employees, "Pedro")
	fmt.Println(employees)
}
