/*
Una empresa de inteligencia artificial necesita tener una funcionalidad para crear una estructura que represente una matriz de datos.
Para ello requieren una estructura Matrix que tenga los métodos:
Set:  Recibe una serie de valores de punto flotante e inicializa los valores en la estructura Matrix
Print: Imprime por pantalla la matriz de una formas más visible (Con los saltos de línea entre filas)
La estructura Matrix debe contener los valores de la matriz, la dimensión del alto, la dimensión del ancho, si es cuadrática
y cuál es el valor máximo.
*/
package main

import "fmt"

type Matrix struct {
	Values      [][]float64
	High        int
	Width       int
	IsQuadratic bool
	MaxValue    float64
}

func (m *Matrix) SetValuesMatrix(values [][]float64, high, width int, isQuadratic bool, maxValue float64) {
	if isQuadratic {
		m.High = high
		m.Width = high
	} else {
		m.High = high
		m.Width = width
	}
	// if(len(values)< high){}
	if (len(values[0]) < width) || (len(values[1]) < width) || (len(values[2]) < width) {
		for i := 0; i < len(values); i++ {
			if len(values[i]) < width {
				dif := width - len(values[i])
				for addZero := 0; addZero < dif; addZero++ {
					values[i] = append(values[i], 0)
				}
			}
		}
	}
	m.Values = values
	m.IsQuadratic = isQuadratic
	m.MaxValue = maxValue
}

func (m Matrix) PrintValuesMatrix() {
	for i := 0; i < m.High; i++ {
		for j := 0; j < m.Width; j++ {
			fmt.Printf("%10.2f", m.Values[i][j])
		}
		fmt.Println("")
	}
}

func main() {
	matrix1 := Matrix{}
	matrix1.SetValuesMatrix(
		[][]float64{
			{20.1},
			{20.5, 0, 90.5, 0, 0},
			{},
			{20.9, 11.9, 90.9, 0, 0, 0},
		}, 3, 4, false, 9)

	matrix1.PrintValuesMatrix()
}
