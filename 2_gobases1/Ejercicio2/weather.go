package main

import "fmt"

func main() {
	var temp int = 25                  // La temperatura se puede declarar como un entero
	var humidity int = 27              // La humedad se puede declarar como un entero
	var atmosphericPressure int = 1011 // La presión atmosférica se puede declarar como un entero

	fmt.Println("La temperatura es:", temp, "º")
	fmt.Println("La humedad es:", humidity, "%")
	fmt.Println("La presión atmosférica es:", atmosphericPressure)
}
