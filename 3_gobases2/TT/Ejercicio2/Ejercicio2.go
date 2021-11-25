package main

import (
	"fmt"
	"math"
)

func roundToInt(x float64) int {
	return int(math.Round(x))
}

type Matrix struct {
	Values    []float64 `json:"matrix_values"`
	Alto      int       `json:"matrix_alto"`
	Ancho     int       `json:"matrix_ancho"`
	Is_Square bool      `json:"matrix_is_square"`
	Max_Value float64   `json:"matrix_max_value"`
}

func (m *Matrix) Set(values ...float64) {
	var aux float64 = float64(len(values))
	var sqrt = math.Sqrt(aux)
	height := roundToInt(sqrt)
	div := (float64)(aux / (float64)(height))
	width := roundToInt(div)

	if (float64)(width*height) < aux {
		width++
	}

	var max float64 = -999999999.0
	for _, element := range values {
		if element > max {
			max = element
		}
	}

	m.Values = values
	m.Alto = height
	m.Ancho = width
	m.Is_Square = (height == width) && (aux == (float64)(height*width))
	m.Max_Value = max
}

func (m Matrix) Print() {
	var a, b int = 0, m.Ancho
	for i := 1; i < m.Alto; i++ {
		fmt.Printf("Fila %d: %v\n", i, m.Values[a:b])
		a = b
		b = b + m.Ancho
	}
	b = len(m.Values)
	fmt.Printf("Fila %d: %v\n", m.Alto, m.Values[a:b])
}

func main() {
	var m Matrix
	m.Set(4.6, 3, 4, 5, 6.9, 156, 34, 2, 89, 64, 7, 21, 56, 43, 21, 999)
	m.Print()
	fmt.Printf("%+v\n", m)
}
