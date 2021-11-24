package main

import "fmt"

func calcularPromedio(calificaciones ...float64) float64 {
	cantidad := len(calificaciones)
	var suma float64

	for i := 0; i < cantidad; i++ {
		suma = suma + calificaciones[i]
	}
	return suma / float64(cantidad)
}

func main() {

	promedio := calcularPromedio(10, 10, 2, 5.50, 10, 7, 10)
	fmt.Printf("El promedio es %.2f \n", promedio)
}
