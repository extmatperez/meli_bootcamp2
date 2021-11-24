package main

import "fmt"

func main() {
	estudiantes := []string{"Benjamin", "Nahuel", "Brenda", "Marcos", "Pedro", "Axel", "Alez", "Dolores", "Federico", "Hernán", "Leandro", "Eduardo", "Duvraschka"}
	fmt.Println(estudiantes)
	fmt.Println("Ingrese el nuevo estudiante:")
	var estudiante string
	fmt.Scanf("%v", &estudiante)
	estudiantes = append(estudiantes, estudiante)
	fmt.Println(estudiantes)
}
