package main

import "fmt"

func main() {
	var temperatura int
	var humedad float64
	var presion float64

	temperatura = 17
	humedad = 78
	presion = 21.67

	fmt.Println("la temperatura en Bogotá es de:", temperatura)
	fmt.Println("la humedad del ambiente es de:", humedad, "%")
	fmt.Println("la presion atmosferica es de:", presion)

	/*
		Para la temperatura asigne un entero porque
		mayormente se usa ese tipo de valor en la temperatura.
		La humedad la trabajo en float porque puede contener valores decimales
		y al final la concateno con un porcentaje porque asi esa el estandar de
		la humedad
		La presion atmosferica tambien la trabajo en float porque casi siempre
		contiene valores con punto flotante
	*/
}
