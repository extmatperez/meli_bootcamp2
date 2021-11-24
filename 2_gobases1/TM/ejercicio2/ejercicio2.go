package main

import "fmt"

func main() {
	var temp float32
	var humidity float32
	var pressure float32

	temp = 23
	humidity = 36
	pressure = 1020
	fmt.Printf("Temperatura: %2.2f C", temp)
	fmt.Printf("Humedad: %2.2f %%", humidity)
	fmt.Printf("Presion Atmosferica: %2.2f Hpa", pressure)
}
