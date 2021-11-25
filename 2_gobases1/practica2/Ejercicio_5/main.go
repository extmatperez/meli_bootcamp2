package main

import "fmt"

func main() {

	var students []string = []string{"Benjamin", "Nahuel", "Brenda", "Marcos", "Pedro", "Axel",
						 "Alez", "Dolores", "Federico", "Hernán", "Leandro", "Eduardo", "Duvraschka"}
	fmt.Println("Los estudiantes son: ", students)
	fmt.Println("Se agrega a Gabriela")
	students = append(
		students,
		"Gabriela",
	)
	fmt.Println("nueva lista de estudiantes: ", students)
} 