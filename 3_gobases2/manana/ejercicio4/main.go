package main

import (
	"fmt"
)

const (
	minimo   = "minimo"
	promedio = "promedio"
	maximo   = "maximo"
)

func main() {

	//minFunc, err := operacion(minimo)
	//promFunc, err := operacion(promedio)
	//maxFunc, err := operacion(maximo)

	//valorMinimo := minFunc(2,3,3,4,1,2,4,5)
	//valorPromedio := promFunc(2,3,3,4,1,2,4,5)
	//valorMaximo := maxFunc(2,3,3,4,1,2,4,5)

	fmt.Println("Hola mundo")
}

func operacion(calculo string) func(calculo string, valor ...float64) {
	switch calculo {
	case minimo:

	case promedio:

	case maximo:

	default:

	}
	return nil
}

func calcular(calculo string, valor ...float64) {
	switch calculo {
	case minimo:

	case promedio:

	case maximo:

	default:

	}
}

func minFunc(valores ...float64) float64 {
	numero := 0.0
	for _, valor := range valores {
		if numero < valor {
			numero = valor
		}
	}
	return numero
}

func maxFunc(valores ...float64) float64 {
	numero := 0.0
	for _, valor := range valores {
		if numero > valor {
			numero = valor
		}
	}
	return numero
}

func promFunc(valores ...float64) float64 {
	sumatoria := 0.0
	for _, valor := range valores {
		sumatoria += valor

	}
	promedio := sumatoria / float64(len(valores))
	return promedio
}
