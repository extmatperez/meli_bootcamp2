package main

import "fmt"

func main() {
	estudiantes := []string{"Benjamin", "Nahuel", "Brenda", "Marcos",
		"Pedro", "Axel", "Alez", "Dolores",
		"Federico", "Hernán", "Leandro", "Eduardo", "Duvraschka", "Gabriela"}
	for _, s := range estudiantes {
		fmt.Printf("%s\n", s)
	}
}
