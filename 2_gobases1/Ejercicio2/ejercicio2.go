package main

import "fmt"

func main() {

	var (
		temperatura = 30.0
		humedad     = 65
		presion     = 1012
	)
	//fmt.Println("La Humedad es:", humedad, "%", "%T", humedad)
	fmt.Printf("La temprera es de: %v ºC y el tipo de dato es %T \n", temperatura, temperatura)
	fmt.Printf("La humedad relativa es de: %v %% y el tipo de dato es %T \n", humedad, humedad)
	fmt.Printf("La Presión es de: %vhPa y el tipo de dato es %T \n", presion, presion)
}
