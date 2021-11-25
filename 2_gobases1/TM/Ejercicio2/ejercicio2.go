package main

import "fmt"

func main() {
	var (
		temperatura int     = 19
		humedad     int     = 54
		presion     float64 = 1019.6
	)

	fmt.Printf("La temperatura en Montevideo en grados centigrados es: %d\n", temperatura)
	fmt.Printf("La humedad es de: %d%%\n", humedad)
	fmt.Printf("Y la presion atmosferica es: %.2f", presion)
}
