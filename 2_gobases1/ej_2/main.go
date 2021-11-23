package main

import "fmt"

var temperature, humidity, pressure []float64

func main() {
	temperature = []float64{10.2, 20.3, 30.4, 40.5, 50.6}
	humidity = []float64{10.2, 20.3, 30.4, 40.5, 50.6}
	pressure = []float64{10.2, 20.3, 30.4, 40.5, 50.6}
	fmt.Println("Temperaturas: ", temperature)
	fmt.Println("Humedades: ", humidity)
	fmt.Println("Presiones: ", pressure)
}