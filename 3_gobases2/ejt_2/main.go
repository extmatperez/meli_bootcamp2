package main

import "fmt"

type Matrix struct {
	alto int
	ancho int
	valores [][]float64
	isCuadrada bool
	max float64
}

func (m *Matrix) Set (valores [][]float64) {
	m.valores = valores
	m.alto = len(valores)
	m.ancho = len(valores[0])
	m.isCuadrada = m.alto == m.ancho
	m.max = 0
	for i := 0; i < m.alto; i++ {
		for j := 0; j < m.ancho; j++ {
			if m.max < m.valores[i][j] {
				m.max = m.valores[i][j]
			}
		}
	}
}

func (m *Matrix) Print() {
	for i := 0; i < m.alto; i++ {
		for j := 0; j < m.ancho; j++ {
			fmt.Printf("%.2f\t", m.valores[i][j])
		}
		fmt.Println()
	}
}

func main() {
	m := Matrix{}
	m.Set([][]float64{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	})
	m.Print()
}