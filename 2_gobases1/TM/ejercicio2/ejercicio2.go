/*
Una empresa de meteorología quiere tener una aplicación donde pueda tener la temperatura y humedad y
presión atmosférica de distintos lugares.
Declara 3 variables especificando el tipo de dato, como valor deben tener la temperatura, humedad y
presión de donde te encuentres.
Imprime los valores de las variables en consola.
¿Qué tipo de dato le asignarías a las variables?
Para la variable de temperatura me pareció mejor idea usar float y para porcentaje 
*/

package main

import "fmt"

func main() {
	var temperatura, humedad float64
	var presionAtmosferica int

	temperatura = 28.6
	humedad = 56.9
	presionAtmosferica = 1018

	fmt.Println("La temperatura actual en Córdoba es de", temperatura, "ºC, el procentaje de humedad es del", humedad, "% y la presión atmosférica es", presionAtmosferica,"hPa")
}