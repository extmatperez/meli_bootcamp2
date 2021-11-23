package main

import "fmt"

func main() {
	/*var (
		temperature = 26.3
		humidity = 64
		pressure = 1.45
	)*/

	var temperature, humidity, pressure float64 = 26.3, 64, 992.4

	// Esta buena la aclaracion de los %
	fmt.Printf("En este momento en Córdoba hace %v grados, una humedad de %v%% y una presión atmosférica de %vhPa.", temperature, humidity, pressure)
}
