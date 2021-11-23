package main

import (
	"fmt"
)

func main() {
	var temperatura, humedad, presion float64 = 19, 59, 1020
	fmt.Printf("La temperatura es: %.0fºC , la humedad es del %.2f%% y la presión de %.0f milibares", temperatura, humedad, presion)

}
