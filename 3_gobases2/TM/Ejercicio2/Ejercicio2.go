package main

import (
	"errors"
	"fmt"
)

func orquestadorOperaciones(valores []float64) float64 {
	var resultado float64
	for i, valor := range valores {
		if i == 0 {
			resultado = valor
		} else {
			resultado = opSuma(resultado, valor)
		}
	}
	return resultado
}

func opSuma(valor1, valor2 float64) float64 {
	return valor1 + valor2
}

func mean_calc(qual ...float64) (float64, error) {

	for _, element := range qual {
		if element < 0 {
			return 0, errors.New("No puede haber nota negativa.")
		}
	}
	var mean float64
	qual_sum := orquestadorOperaciones(qual)
	mean = qual_sum / (float64)(len(qual))
	return mean, nil
}

func main() {
	calculated_mean, err := mean_calc(3, 8, 9, 8, 10)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("El promedio es %1.1f", calculated_mean)
	}
}
