package ej6

import "fmt"

var employees = map[string]int{
	"Benjamin": 20,
	"Nahuel":   26,
	"Brenda":   19,
	"Dario":    44,
	"Pedro":    30,
}

func Ej6() int {
	fmt.Println(employees["Benjamin"])
	return employees["Benjamin"]
}

func SearchEmployee(employee string) int {
	fmt.Println(employees[employee])
	fmt.Println(employee)

	return employees[employee]
}

func GreaterThan21() int {
	counter := 0

	for _, val := range employees {
		if val > 21 {
			counter++
		}
	}
	fmt.Printf("Employees with age greater than 21: %d \n", counter)
	return counter
}

func AddEmployee(name string, age int) map[string]int {
	employees[name] = age
	fmt.Println(employees)
	return employees
}

func DeleteEmployee(name string) map[string]int {
	delete(employees, name)
	fmt.Println(employees)
	return employees
}
