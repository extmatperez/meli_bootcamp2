package main

import "fmt"

func main() {
	var estudiantes = []string{"Benjamin", "Nahuel", "Brenda", "Marcos", "Pedro", "Axel", "Alez", "Dolores", "Federico", "Hernán", "Leandro", "Eduardo", "Duvraschka"}

	fmt.Printf("Estudiantes: %v \n", estudiantes)

	estudiantes = append(estudiantes, "Gabriela")

	fmt.Printf("Estudiantes: %v \n", estudiantes)

}
