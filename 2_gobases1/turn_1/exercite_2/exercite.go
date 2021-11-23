package main

import "fmt"

func main() {

	var temperature int
	var humidity float64
	var pressure float64

	temperature = 22
	humidity = 42.1111
	pressure = 730.00

	fmt.Printf("Temperatuere: %d ºC", temperature)
	fmt.Printf("\nHumidity: %.2f %%", humidity)
	fmt.Printf("\nPressure: %v mmHg", pressure)

}
