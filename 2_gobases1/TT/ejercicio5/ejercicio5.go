package main

import "fmt"

func main() {
	students := []string{"Benjamin", "Nahuel", "Brenda", "Marcos", "Pedro", "Axel", "Alez", "Dolores", "Federico", "Hernán", "Leandro", "Eduardo", "Duvraschka"}
	students = append(students, "Brenda")
	for _, student := range students {
		fmt.Printf("Student:%s \n", student)
	}
}
