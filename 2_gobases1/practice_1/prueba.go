package main

import "fmt"

func main() {

	fmt.Println("Hola Mundo")

	myFirstValue := "Digneli"
	fmt.Println("My name is", myFirstValue)

	hours := 65.1

	fmt.Printf("Las \" horas son: %08.2f, %.2v, %T", hours, hours, hours)

	hours = 555.5

	fmt.Printf("Las \" horas son: %08.2f, %v, %T", hours, hours, hours)

	hours = 555555.5

	fmt.Printf("Las \" horas son: %8.2f, %v, %T", hours, hours, hours)

}
