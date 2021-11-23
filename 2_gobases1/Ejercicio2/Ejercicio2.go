package main

import "fmt"

func main() {
	var temperatura, humedad int = 29, 32

	var presion string = "1021 hPa"

	fmt.Printf("Temperatura: %d", temperatura)
	fmt.Printf(" Cº y su tipado es: %T \n", temperatura)

	fmt.Printf("Humedad: %d", humedad)
	fmt.Printf(" y su tipado es : %T \n", humedad)

	fmt.Printf("Presion: %s", presion)
	fmt.Printf(" y su tipado es : %T \n", presion)

}
