package main

import "fmt"

const (
	Suma   = "+"
	Resta  = "-"
	Multip = "*"
	Divis  = "/"
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

func opDivis(valor1, valor2 float64) float64 {
	if valor2 == 0 {
		return 0
	}
	return valor1 / valor2
}

func operacionAritmetica(operador string, valores ...float64) float64 {
	switch operador {
	case Suma:
		return orquestadorOperaciones(valores, opSuma)
	case Resta:
		return orquestadorOperaciones(valores, opResta)
	case Multip:
		return orquestadorOperaciones(valores, opMultip)
	case Divis:
		return orquestadorOperaciones(valores, opDivis)
	}
	return 0
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

func suma(values ...float64) float64 {
	var resultado float64
	for _, v := range values {
		resultado += v
	}
	return resultado

}

func main() {
	fmt.Println(operacionAritmetica(Resta, 1, 2, 3, 4, 5))
}
