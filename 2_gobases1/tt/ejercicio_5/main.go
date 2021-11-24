package main

import "fmt"

func main() {

	var alumnos = []string{
		"Benjamin", "Nahuel", "Brenda", "Marcos", "Pedro",
		"Axel", "Alez", "Dolores", "Federico", "Hernán",
		"Leandro", "Eduardo", "Duvraschka"}

	fmt.Println(alumnos)

	alumnos = append(alumnos, "Gabriela")
	fmt.Println(alumnos)

}
