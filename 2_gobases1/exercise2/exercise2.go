// Exercise Clima
package main

import "fmt"

var temperature float64
var humidity float64
var press float64

func main() {
	temperature = 28
	humidity = 26.5
	press = 1013
	fmt.Printf("En Santiago de Chile se tiene que: \nTemperatura: %f ºC \nPresion %f hPa \nHumedad: %f", temperature, press, humidity)
}
