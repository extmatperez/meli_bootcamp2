package main

import "fmt"

type matrix struct {
	valores     []float64
	alto        int
	ancho       int
	cuadratica  bool
	valorMaximo int
}

func (m *matrix) setValores(alto int, ancho int, cuadratica bool, valores ...float64) {

	m.alto = alto
	m.ancho = ancho
	m.cuadratica = cuadratica
	m.valorMaximo = alto * ancho
	m.valores = valores

}

/* func (m matrix) imprimirMatrix() {
	if m.cuadratica == true {
	}
} */

func (m *matrix) imprimirMatrix() {
	fmt.Println("- - - - - - - - - - - - - - - - - - - - - - - - ")
	fmt.Println("- - - - -Se imprimira la matrix")
	fmt.Println("- - - - - - - - - - - - - - - - - - - - - - - - ")
	var alturaMatrix int
	var anchoMatrix int
	for _, val := range m.valores {
		if anchoMatrix == m.ancho {
			fmt.Println()
			anchoMatrix = 0
			alturaMatrix++
		}
		if alturaMatrix == m.alto {
			break
		}
		fmt.Printf("%.1f\t", val)
		anchoMatrix++
	}
}
func main() {
	matrix1 := matrix{}
	matrix1.setValores(5, 5, true, 11, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 2, 3, 4, 5, 6, 7, 8, 9, 10, 1, 1, 1, 1, 1)
	matrix1.imprimirMatrix()
	fmt.Println()
}
