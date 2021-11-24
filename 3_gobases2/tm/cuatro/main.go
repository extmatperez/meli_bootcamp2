package main

import (
	"errors"
	"fmt"
)

const (
	minimo   = "minimo"
	promedio = "promedio"
	maximo   = "maximo"
)

func operacion(tipoCalculo string) (func(notas ...int) int, error) {
	err := errors.New("calculo no definido")
	switch tipoCalculo {
	case minimo:
		return func(notas ...int) int {
			//valor minimo
			return 0
		}, nil
	case promedio:
		return func(notas ...int) int {
			//valor promedio
			return 0
		}, nil
	case maximo:
		return func(notas ...int) int {
			//valor maximo
			return 0
		}, nil
	default:
		return func(notas ...int) int { return 0 }, err
	}

}

func main() {

	minFunc, _ := operacion(minimo)
	promFunc, _ := operacion(promedio)
	maxFunc, _ := operacion(maximo)
	_, err := operacion("hola")

	valorMinimo := minFunc(7, 6, 8, 9, 7)
	valorPromedio := promFunc(7, 6, 8, 7, 5)
	valorMaximo := maxFunc(7, 10, 8, 9, 8)

	fmt.Println(valorMinimo)
	fmt.Println(valorPromedio)
	fmt.Println(valorMaximo)

	if err == nil {
		fmt.Println(err)
	}
}
