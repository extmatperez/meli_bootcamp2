package main

import "fmt"

func main()  {
	
	fmt.Println("Bienvenido al ejercicio 5!!!")

	var alumnos = []string{"Benjamin", "Nahuel", "Brenda", "Marcos", "Pedro", "Axel", "Alez", "Dolores", "Federico", "Hernán", "Leandro", "Eduardo", "Duvraschka"}

	fmt.Println("Tenemos", len(alumnos), "alumnos y son", alumnos)

	alumnos = append(alumnos, "Gabriela")

	fmt.Println("Ahora tenemos", len(alumnos), "alumnos y son", alumnos)

}