package main

import "fmt"

func main() {

	/* Ejercicio 1 */
	nombre_1 := "Matias"
	apellido_1 := "De Bonis"

	fmt.Println(nombre_1,apellido_1)

	/* Ejercicio 2 */
	var temperatura float64
	temperatura = 19.0
	var humedad float64
	humedad = 35.0
	var presion int
	presion = 1024

	fmt.Printf("Temperatura: %v, Humedad %v%%, presión: %v", temperatura, humedad, presion)
  
}