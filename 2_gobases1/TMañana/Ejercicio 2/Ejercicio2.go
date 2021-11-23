package main

import "fmt"

func main(){
	var temperatura float64
	var humedad int
	var presion int

	temperatura = 31.2
	humedad = 51
	presion = 1022

	fmt.Println("La temperatura es", temperatura, "°C con una humedad del ", humedad,"%", "y la presion es", presion, "hpa")
}