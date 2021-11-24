package main

import "fmt"

const (
	min = "min"
	avg = "avg"
	max = "max"
)

func minFunc(datos ...float64) float64 {
	valor := datos[0]
	for _, value := range datos {
		if value < valor {
			valor = value
		}
	}
	return valor
}

func maxFunc(datos ...float64) float64 {
	valor := datos[0]
	for _, value := range datos {
		if value > valor {
			valor = value
		}
	}
	return valor
}

func promFunc(datos ...float64) float64 {
	prom := 0.0
	suma := 0.0

	for _, value := range datos {
		suma += value
	}

	prom = suma / float64(len(datos))

	return prom
}

func main() {

	prom := operacion(avg)
	r := prom(24, 8, 0, 5, 1)
	fmt.Println(r)

	mini := operacion(min)
	result := mini(24, 8, 0, 5, 1)
	fmt.Println(result)

	oper := operacion(max)
	maxi := oper(24, 8, 0, 5, 1)
	fmt.Println(maxi)

}

func operacion(operador string) func(valores ...float64) float64 {
	switch operador {
	case "min":
		return minFunc
	case "max":
		return maxFunc
	case "avg":
		return promFunc
	}
	return nil
}
