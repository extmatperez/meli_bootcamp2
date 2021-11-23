package main

import "fmt"

func main() {

	var humedad int
	var temperatura int
	var presion float32

	humedad = 47
	presion = 30.06
	temperatura = 22

	fmt.Printf("Humedad: %d%%, presion: %.2f in., temperatura: %dºC\n",
		humedad,
		presion,
		temperatura)
}
