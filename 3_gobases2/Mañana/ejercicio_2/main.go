package main

import (
	"errors"
	"fmt"
)

func promedio(notas ...float64) (float64, error) {
	var prom float64
	suma := 0.00
	contador := 0
	for _, numero := range notas {
		if numero < 0 {
			return 0, errors.New("Valor negativo")
		}
		suma += numero
		contador++
	}
	prom = suma / float64(contador)
	return prom, nil
}

func main() {

	fmt.Println("Notas : 4.5, 4.0, 3.5")
	fmt.Println("Promedio :")
	fmt.Println(promedio(4.3, 4.1, 3.2))
}
