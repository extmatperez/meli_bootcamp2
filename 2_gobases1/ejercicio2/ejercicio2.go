package main

import "fmt"

func main() {
	var temperatura int = 18
	var humedad int = 58
	var presion float32 = 1019

	fmt.Printf("El tiempo en Montevideo es:\n")
	fmt.Printf("Temperatura %d°C. Humedad: %d%%. Presion: %.1f hPa:", temperatura, humedad, presion)

}
