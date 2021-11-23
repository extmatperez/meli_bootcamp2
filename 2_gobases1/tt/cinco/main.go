package main

import "fmt"

var students = []string{"Benjamin", "Nahuel", "Brenda", "Marcos", "Pedro", "Axel", "Alez",
	"Dolores", "Federico", "Hernán", "Leandro", "Eduardo", "Duvraschka",
}

func main() {
	students = append(students, "Gabriela")
	fmt.Println(students)
}
