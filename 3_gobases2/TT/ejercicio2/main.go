package main

import "fmt"

type Matrix struct {
	Alto         int
	Ancho        int
	EsCuadratica bool
	ValorMaximo  float64
	Valores      []float64
}

func main() {
	matriz := Matrix{}
	matriz.Set(3, 4, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12)
	matriz.Print()
}

func (m *Matrix) Set(alto int, ancho int, valores ...float64) {
	m.Alto = alto
	m.Ancho = ancho
	m.EsCuadratica = alto == ancho
	max := valores[0]
	for _, valor := range valores {
		if max < valor {
			max = valor
		}
	}
	m.ValorMaximo = max
	m.Valores = valores
}

func (m *Matrix) Print() {
	for i := 0; i < m.Alto; i++ {
		for j := 0; j < m.Ancho; j++ {
			fmt.Printf("%.2f\t", m.Valores[i*m.Ancho+j])
		}
		fmt.Printf("\n")
	}
}
