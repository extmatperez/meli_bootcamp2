package main

import "fmt"

func main() {

	var (
		temperatura = 30.0
		humedad     = 65
		presion     = 1012
	)
	//fmt.Println("La Humedad es:", humedad, "%", "%T", humedad)
	fmt.Println("")
	fmt.Printf("La temprera es de: %v ºC        -- tipo de dato (%T )\n", temperatura, temperatura)
	fmt.Printf("La humedad relativa es de: %v %% -- tipo de dato (%T) \n", humedad, humedad)
	fmt.Printf("La Presión es de: %vhPa       -- tipo de dato (%T) \n", presion, presion)
	fmt.Println("")
}
