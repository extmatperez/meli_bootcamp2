package main

import "fmt"

func main() {
	var temperatura float64
	temperatura = 19.0
	var humedad float64
	humedad = 35
	var presion float64
	presion = 1024.0

	fmt.Printf("Temperatura: %v, Humedad %v%%, presión: %v", temperatura,humedad, presion)
}