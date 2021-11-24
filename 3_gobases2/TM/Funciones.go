package main

import (
	"errors"
	"fmt"
)

const (
	Suma     = "+"
	Resta    = "-"
	Multip   = "*"
	Division = "/"
)

func opSuma(valor1, valor2 float64) float64 {
	return valor1 + valor2
}

func opResta(valor1, valor2 float64) float64 {
	return valor1 - valor2
}

func opMultip(valor1, valor2 float64) float64 {
	return valor1 * valor2
}

func opDivision(valor1, valor2 float64) float64 {
	if valor2 == 0 {
		return 0
	}
	return valor1 / valor2
}

func orquestadorOperaciones(valores []float64, operacion func(value1, value2 float64) float64) float64 {
	var resultado float64
	for i, valor := range valores {
		if i == 0 {
			resultado = valor
		} else {
			resultado = operacion(resultado, valor)
		}
	}
	return resultado
}

// Tenemos las ellipsis que son un conjunto de valores que se llaman con, por ejemplo, si pueden ser muchos float, con ...float64.
// En este caso esta es una funcion que recibe la operacion a realizar y los valores que se van a aplicar.
func operacionAritmetica(operador string, valores ...float64) float64 {
	switch operador {
	case Suma:
		return orquestadorOperaciones(valores, opSuma)
	case Resta:
		return orquestadorOperaciones(valores, opResta)
	case Multip:
		return orquestadorOperaciones(valores, opMultip)
	case Division:
		return orquestadorOperaciones(valores, opDivision)
	}
	return 0
}

// Funcion de division con gestion y creacion de errores
func division(dividendo, divisor float64) (float64, error) {
	if divisor == 0 {
		return 0, errors.New("El divisor no puede ser cero.")
	}
	return dividendo / divisor, nil
}

func main() {

	fmt.Println(operacionAritmetica(Suma, 2, 4, 5, 6, 1, 1, 1, 2))

	res, err := division(2, 2)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("La división es ", res)
	}
}
