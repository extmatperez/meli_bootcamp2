package main

import "fmt"

func main() {
	var matrix Matrix

	fmt.Println("==========Matriz inicial==========")
	matrix.printMatrixValues()
	fmt.Println("==================================")
	fmt.Println("")

	matrix.setValues(
		[][]float64{
			{1, 2, 3},
			{4, 5, 6},
			{7, 8, 9},
		}, 3, 3, true, 9)

	fmt.Println("===========Matriz final===========")
	matrix.printMatrixValues()
	fmt.Println("==================================")
}

type Matrix struct {
	Values      [][]float64
	High        float64
	Width       float64
	IsQuadratic bool
	MaxValue    float64
}

func (m *Matrix) setValues(values [][]float64, high, width float64, isQuadratic bool, maxValue float64) {
	m.Values = values
	m.High = high
	m.Width = width
	m.IsQuadratic = isQuadratic
	m.MaxValue = maxValue
}

func (m Matrix) printMatrixValues() {
	for i := 0; i < len(m.Values); i++ {
		for j := 0; j < len(m.Values[i]); j++ {
			fmt.Printf("%.2f ", m.Values[i][j])
		}
		fmt.Println("")
	}
	fmt.Println("Alto:", m.High)
	fmt.Println("Ancho:", m.Width)
	fmt.Println("Es cuadratica:", m.IsQuadratic)
	fmt.Println("Valor máximo:", m.MaxValue)
}
