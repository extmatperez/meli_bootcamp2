package main

import "fmt"

func main() {
	var matrix Matrix
	matrix.Set(2, 3, 5, 7, 12, 78, 34, 4, 45, 45)
	matrix.Print()
}

type Matrix struct {
	Altura       uint16
	Ancho        uint16
	EsCuadratica bool
	ValorMaximo  uint16
	Valores      []float64
}

func (m *Matrix) Set(altura uint16, ancho uint16, valores ...float64) {
	m.Altura = altura
	m.Ancho = ancho
	if m.Altura == m.Ancho {
		m.EsCuadratica = true
	}
	m.ValorMaximo = m.Altura * m.Ancho
	m.Valores = valores
}

func (m *Matrix) Print() {
	var altura uint16
	var anchura uint16
	for _, val := range m.Valores {
		//fmt.Println(altura)
		//fmt.Println(anchura)
		if anchura == m.Ancho {
			fmt.Printf("\n")
			anchura = 0
			altura++
		}
		if altura == m.Altura {
			break
		}
		fmt.Printf("%.1f\t", val)
		anchura++
	}
}
