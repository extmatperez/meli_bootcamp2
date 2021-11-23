package main

import "fmt"

func main() {
	var temperatura float32
	var humedad float32
	var presion float32

	temperatura = 17
	humedad = 67.2
	presion = 548

	fmt.Printf("Temperatura: %.2fº", temperatura)
	fmt.Printf("\nHumedad: %.2f%%", humedad)
	fmt.Printf("\nPresión: %.2f mmHg", presion)
}

/*
3. Pienso que pueden ser válidos valores de float o int, dependiendo de la exactitud que se requiera.
*/
