package main

import "fmt"

const (
	Suma     = "+"
	minimo   = "minimo"
	maximo   = "maximo"
	promedio = "promedio"
)

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

func opSuma(valor1, valor2 float64) float64 {
	return valor1 + valor2
}

func opMaximo(valores []float64) float64 {
	var max float64
	for i, element := range valores {
		if i == 0 {
			max = element
		} else {
			if element > max {
				max = element
			}
		}
	}
	return max
}

func opMinimo(valores []float64) float64 {
	var min float64
	for i, element := range valores {
		if i == 0 {
			min = element
		} else {
			if min > element {
				min = element
			}
		}
	}
	return min
}

func opPromedio(valores []float64) float64 {
	suma := orquestadorOperaciones(valores, opSuma)
	return suma / (float64)(len(valores))
}

func operacionAritmetica(operador string, valores ...float64) float64 {
	switch operador {
	case Suma:
		return orquestadorOperaciones(valores, opSuma)
	case minimo:
		return opMinimo(valores)
	case maximo:
		return opMaximo(valores)
	case promedio:
		return opPromedio(valores)
	}
	return 0
}

func main() {
	fmt.Println(operacionAritmetica(promedio, 2, 4, 5, 6, 1, 1, 1, 2))
}
