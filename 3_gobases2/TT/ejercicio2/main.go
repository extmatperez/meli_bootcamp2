package main

import "fmt"

type Matrix struct {
	Valores    [][]float64
	Ancho      int
	Alto       int
	EsCuadrada bool
	Max        float64
}

func (m *Matrix) set(vals ...float64) {
	if m.Alto == m.Ancho {
		m.EsCuadrada = true
	}
	m.Valores = make([][]float64, m.Alto)
	for i := 0; i < m.Alto; i++ {
		var nuevaFila []float64
		for j := 0; j < m.Ancho; j++ {
			nuevaFila = append(nuevaFila, vals[j+m.Ancho*i])
		}
		m.Valores[i] = append(m.Valores[i], nuevaFila...)
	}
}

func (m *Matrix) print() {
	for _, fila := range m.Valores {
		fmt.Println(fila)
	}
}

func main() {
	matrix1 := Matrix{Ancho: 3, Alto: 3}
	matrix1.set(1, 2, 3, 4, 5, 6, 7, 8, 9)
	fmt.Println(matrix1)
	matrix1.print()

}
