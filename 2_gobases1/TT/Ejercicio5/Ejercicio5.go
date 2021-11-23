package main

import "fmt"

func main() {
	var students = []string{"Benjamin", "Nahuel", "Brenda", "Marcos", "Pedro", "Axel", "Alez", "Dolores", "Federico", "Hernán",
		"Leandro", "Eduardo", "Duvraschka"}
	fmt.Printf("Listado de estudiantes: %v\n", students)
	students = append(students, "Gabriela")
	fmt.Printf("Listado de estudiantes: %v\n", students)
}
