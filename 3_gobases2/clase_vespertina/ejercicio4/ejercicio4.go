package main

import (
	"errors"
	"fmt"
)

func main() {
	notas := leerNotas()
	funcion := obtenerFuncion("maximo")
	val, err := funcion(notas)
	if err == nil {
		fmt.Printf("El resultado es: %v \n", val)
	} else {
		fmt.Println("Hubo un error")
	}
}

func leerNotas() []float64 {
	var notas []float64
	fmt.Println("Para salir escriba -1 \n")
	for {
		nota := capturarNotas()
		if nota != -1 {
			notas = append(notas, nota)
		} else {
			break
		}
		fmt.Printf("\n")
	}
	return notas
}

func capturarNotas() float64 {
	var nota float64
	fmt.Println("Ingrese la nota:")
	fmt.Scanf("%f", &nota)
	return nota
}

func obtenerFuncion(operacion string) func(notas []float64) (float64, error) {
	switch operacion {
	case "minimo":
		return func(notas []float64) (float64, error) {
			if len(notas) == 0 {
				return 0, errors.New("No hay notas existentes")
			}
			aux := 0.0
			for _, val := range notas {
				if val < aux || aux == 0 {
					aux = val
				}
			}
			return aux, nil
		}

	case "maximo":
		return func(notas []float64) (float64, error) {
			if len(notas) == 0 {
				return 0, errors.New("No hay notas existentes")
			}
			aux := 0.0
			for _, val := range notas {
				if val > aux || aux == 0 {
					aux = val
				}
			}
			return aux, nil
		}

	case "promedio":
		return func(notas []float64) (float64, error) {
			if len(notas) == 0 {
				return 0, errors.New("No hay notas existentes")
			}
			total := 0.0
			for _, val := range notas {
				total += val
			}
			return (total / float64(len(notas))), nil
		}
	default:
		return nil
	}
}
