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

func operacion(operation string) (int, error) {

	switch operation {
	case minimo:
		return minFunc(2, 3, 3, 4, 1, 2, 4, 5), nil
	case promedio:
		return promFunc(2, 3, 3, 4, 1, 2, 4, 5), nil
	case maximo:
		return maxFunc(2, 3, 3, 4, 1, 2, 4, 5), nil
	default:
		return 0, errors.New("Operador inválido")
	}

}

func minFunc(num ...int) {

}

func main() {

	resultado, err := operacion("minimo")

	if err != nil {
		fmt.Println("error")
	}

}
