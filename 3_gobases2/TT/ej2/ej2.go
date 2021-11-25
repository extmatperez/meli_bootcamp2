package ej2

import (
	"fmt"
)

type Matrix struct {
	Values          [][]float64
	HighDimention   int  `default:"0"`
	WeightDimention int  `default:"0"`
	IsCuadratic     bool `default:"true"`
	MaxValue        float64
}

func Ej2() {
	firstMatrix := [][]float64{{1.0, 2.1, 3.4}}
	// newMatrix := Matrix{firstMatrix, 1, 3, false, 0}
	newMatrix := new(Matrix)

	otherMatrix := []float64{5.5, 3.4, 9.0, 1.1}

	fmt.Println(newMatrix)
	fmt.Println("Print values")
	newMatrix.Print()
	fmt.Println("Add first matrix inside")
	fmt.Println(newMatrix.Set(firstMatrix[0]...))
	fmt.Println("Print matrix after add a new one")
	fmt.Println(newMatrix)
	fmt.Println("Add other matrix inside")
	fmt.Println(newMatrix.Set(otherMatrix...))
	fmt.Println("Printing values again")
	newMatrix.Print()
	fmt.Println("Print matrix after add a new one")
	fmt.Println(newMatrix)
}

func (matrix *Matrix) Set(values ...float64) [][]float64 {
	if matrix.WeightDimention < len(values) {
		matrix.WeightDimention = len(values)
	}

	matrix.HighDimention++

	for _, el := range values {
		if el > matrix.MaxValue {
			matrix.MaxValue = el
		}
	}

	if matrix.WeightDimention%matrix.HighDimention == 0 {
		matrix.IsCuadratic = true
	} else {
		matrix.IsCuadratic = false
	}

	matrix.Values = append(matrix.Values, values)
	return matrix.Values
}

func (matrix Matrix) Print() {
	for i, el := range matrix.Values {
		fmt.Printf("%v", i)
		fmt.Println(el)
	}
}
