package main

import "fmt"

func main() {
	students := []string{"Benjamin", "Nahuel", "Brenda", "Marcos", "Pedro", "Axel", "Alez", "Dolores", "Federico", "Hernán", "Leandro", "Eduardo", "Duvraschka"}
	fmt.Println(students)

	students = append(students, "Gabriela")
	fmt.Println(students)
}
