package main

import (
	"fmt"
)

func main() {
	estudiantes := []string{"Benjamin", "Nahuel", "Brenda", "Marcos", "Pedro", "Axel", "Alez", "Dolores", "Federico", "Hernán", "Leandro", "Eduardo", "Duvraschka"}
	fmt.Println(estudiantes)
	estudiantes = append(estudiantes, "Gabriela")
	fmt.Println("\n", estudiantes)
}
