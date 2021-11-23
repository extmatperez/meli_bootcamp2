package main

import "fmt"

func main() {
	/*var (
		temperature = 26.3
		humidity = 64
		pressure = 1.45
	)*/

	var temperature, humidity, pressure float64 = 26.3, 64, 1.45

	fmt.Println("En este momento en Córdoba hace ", temperature, " grados, una humedad de ", humidity, "%, y una presión atmosférica de ", pressure, "hPa.")
}
