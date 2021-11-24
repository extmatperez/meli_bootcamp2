package main

import "fmt"

func main() {

	var alumnos = []string{"Nahuel", "Brenda", "Marcos", "Pedro", "Axel", "Dolores", "Federico", "Hernan", "Leandro", "Eduardo", "Duvraschka"}
	for i := 0; i < len(alumnos); i++ {
		fmt.Println("Nombre de alumno: ", alumnos[i])
	}
	//Agregamos un alumno mas al slice
	alumnos = append(alumnos, "Gabriela")
	fmt.Println(alumnos)

}
