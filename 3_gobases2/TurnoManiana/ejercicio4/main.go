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

func operacion(operador string) (int, error) {
	switch operador {
	case minimo:
		return minFunc(2, 3, 3, 4, 7, 2, 4, 5), nil
	case promedio:
		return promFunc(2, 3, 3, 4, 7, 2, 4, 5), nil
	case maximo:
		return maxFunc(2, 3, 3, 4, 7, 2, 4, 5), nil
	default:
		return 0, errors.New("Operador inválido")
	}
}

func minFunc(valores ...int) int {
	minimo := valores[0]
	for i := 0; i < len(valores); i++ {
		if valores[i] < minimo {
			minimo = valores[i]
		}
	}
	return minimo
}

func maxFunc(valores ...int) int {
	maximo := valores[0]
	for i := 0; i < len(valores); i++ {
		if valores[i] > maximo {
			maximo = valores[i]
		}
	}
	return maximo
}

func promFunc(valores ...int) int {
	cantidad := len(valores)
	var suma int
	for i := 0; i < cantidad; i++ {
		suma = suma + valores[i]
	}
	return suma / cantidad
}

func main() {

	resultado, err := operacion("maximo")
	if err != nil {
		fmt.Println("Hubo un error el operador")
	} else {
		fmt.Printf("El resultado de la operacion es %v\n", resultado)
	}
}
