package main

import "fmt"

// Listado de nombres
func main() {

	estudiantes := []string{"Benjamin", "Nahuel", "Brenda", "Marcos", "Pedro", "Axel", "Alez", "Dolores", "Federico", "Hernán", "Leandro", "Eduardo", "Duvraschka"}

	fmt.Println(estudiantes)

	estudiantes = append(estudiantes, "Gabriela")

	fmt.Println(estudiantes)

}
