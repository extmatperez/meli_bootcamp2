package main

import "fmt"

func main() {
	estudiantes := make([]string, 1)
	estudiantes = append(estudiantes, "Benjamin")
	estudiantes = append(estudiantes, "Nahuel")
	estudiantes = append(estudiantes, "Brenda")
	estudiantes = append(estudiantes, "Marcos")
	estudiantes = append(estudiantes, "Pedro")
	estudiantes = append(estudiantes, "Axel")
	estudiantes = append(estudiantes, "Alez")
	estudiantes = append(estudiantes, "Dolores")
	estudiantes = append(estudiantes, "Federico")
	estudiantes = append(estudiantes, "Hernán")
	estudiantes = append(estudiantes, "Leandro")
	estudiantes = append(estudiantes, "Eduardo")
	estudiantes = append(estudiantes, "Duvraschka")

	fmt.Println(estudiantes)
}
