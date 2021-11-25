package main

import "fmt"

func promedio(notas ...float64) float64 {

	var prom float64

	for _, nota := range notas {
		prom += nota
	}
	prom = prom / float64(len(notas))
	return prom
}

func main() {
	var notas = []float64{10.0, 7.0, 5.0, 8.0, 9.0, 9.0}
	fmt.Println(promedio(notas...))
}
