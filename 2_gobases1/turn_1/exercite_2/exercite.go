package main

import "fmt"

func main() {

	var temperature int
	var humidity float64
	var pressure int

	temperature = 22
	humidity = 42.1111
	pressure = 730

	fmt.Printf("Temperatuere: %d", temperature)
	fmt.Printf("\nHumidity: %.2f", humidity)
	fmt.Printf("\nPressure: %d", pressure)

}
