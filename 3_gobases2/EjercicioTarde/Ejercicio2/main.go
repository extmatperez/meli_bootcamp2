package main

import "fmt"

type Matrix struct {
	Alto       int
	Ancho      int
	Cuadratica bool
	Maximo     float64
	Valores    []float64
}

func (m *Matrix) set(valores ...float64) {
	m.Valores = valores
}

func (m Matrix) print() {
	max := m.Alto * m.Ancho
	fmt.Println("Matrix Print: ")
	for i := 0; i < max; i++ {
		if i < len(m.Valores) {
			fmt.Printf(" %8.2f ", m.Valores[i])
		} else {
			fmt.Printf(" %8.2f ", 0.0)
		}

		if (i+1)%m.Ancho == 0 {
			fmt.Printf("\n")
		}
	}
	fmt.Printf("\n")
}

func (m *Matrix) esCuadratica() {
	if m.Ancho == m.Alto {
		m.Cuadratica = true
	} else {
		m.Cuadratica = false
	}
}

func (m *Matrix) maximo() {
	max := m.Valores[0]
	for _, value := range m.Valores {
		if value > max {
			max = value
		}
	}
	m.Maximo = max
}

func main() {

	m1 := Matrix{
		Alto:  3,
		Ancho: 2,
	}
	m1.set(14.0, 15.0, 10.0, 1.0, 2.0, 3.0)
	m1.maximo()
	m1.esCuadratica()
	m1.print()
	fmt.Printf("Matrix m1: %+v \n", m1)

	m2 := Matrix{
		Alto:  4,
		Ancho: 4,
	}
	m2.set(20.0, 30.0, 40.0, 20.0, 1.0, 2.0, -1.0, 9.0, -9.0, 40.0, -40.0, 1.0, 3.4, 5.44, 0.111, 1.11111)
	m2.maximo()
	m2.esCuadratica()
	m2.print()
	fmt.Printf("Matrix m2: %+v \n", m2)

	m3 := Matrix{
		Alto:  1,
		Ancho: 1,
	}
	m3.set(20.0)
	m3.maximo()
	m3.esCuadratica()
	m3.print()
	fmt.Printf("Matrix m3: %+v \n", m3)
}
