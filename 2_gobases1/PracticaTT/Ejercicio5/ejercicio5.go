package main

import "fmt"

func main() {
	students := []string{"Benjamin", "Nahuel", "Brenda", "Marcos", "Pedro", "Axel",
		"Alez", "Dolores", "Federico", "Hernán", "Leandro", "Eduardo", "Duvraschka"}

	fmt.Printf("Estudiantes iniciales: %v\n", students)

	students = append(students, "Gabriela")

	fmt.Printf("Estudiantes actualizados: %v\n", students)
}
