package main

import (
	"fmt"
)

func main() {
	
	var	temperatura int = 30
	var humedad float64 = 60.15
	var	presion float64 = 0.9
	
	fmt.Println("La tempreatura del ambiente es: ", temperatura, "ºC")
	fmt.Println("El porcentaje de humedad es: ", humedad, "%")
	fmt.Println("La presion del ambiente es: ", presion , "atm")
}