// Ejercicio 2 - Clima

// Una empresa de meteorología quiere tener una aplicación donde pueda tener la temperatura y humedad y presión atmosférica de distintos lugares. 
// 1- Declara 3 variables especificando el tipo de dato, como valor deben tener la temperatura, humedad y presión de donde te encuentres.
// 2- Imprime los valores de las variables en consola.
// 3-¿Qué tipo de dato le asignarías a las variables?

package main

import "fmt"

func main() {
	var temperatura int = 24
	var humedad int =  5
	var presion float64 = 800.3

	fmt.Println("Temperatura: ", temperatura)
	fmt.Println("Humedad: ", humedad,"%")
	fmt.Println("Presión: ", presion)
}