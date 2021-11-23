package main

import "fmt"

func main() {
	alumnos := []string{"Benjamin", "Nahuel", "Brenda", "Marcos", "Pedro", "Axel", "Alez", "Dolores", "Federico", "Hernán", "Leandro", "Eduardo", "Duvraschka"}
	alumnos = append(alumnos, "Gabriela")
	fmt.Println(alumnos)
}