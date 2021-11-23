package main

import "fmt"

// Una empresa de meteorología quiere tener una aplicación donde pueda tener la temperatura y humedad y presión atmosférica de distintos lugares. 
// Declara 3 variables especificando el tipo de dato, como valor deben tener la temperatura, humedad y presión de donde te encuentres.
// Imprime los valores de las variables en consola.
// ¿Qué tipo de dato le asignarías a las variables?

func main () {

	var temperatura int
	var humedad uint
	var presion float64

	temperatura = 9
	humedad = 53
	presion = 1017.4

	fmt.Printf("La temperatura actual es: %dº \n La humedad es de: %d%% \n La presion: %1f hPa", temperatura, humedad, presion)
	

}