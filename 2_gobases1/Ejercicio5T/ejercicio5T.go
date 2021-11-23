package main

import "fmt"

func main() {

	estudianteSlice := []string{"Benjamin", "Nahuel", "Brenda", "Marcos", "Pedro", "Axel",
		"Alez", "Dolores", "Federico", "Hernán", "Leandro", "Eduardo", "Duvraschka"}

	fmt.Println("Estudiantes iniciales:\n\t", estudianteSlice)
	estudianteSlice = append(estudianteSlice, "Gabriela")
	fmt.Println("Estudiantes luego de dos semanas:\n\t", estudianteSlice)

}
