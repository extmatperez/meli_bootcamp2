package main

import "fmt"

type Matrix struct {
	Alto       float64
	Ancho      float64
	Cuadratica float64
	MaxValor   float64
}

func (m *Matrix) Set(valores ...float64) {
	m.Alto = valores[0]
	m.Ancho = valores[1]
	m.Cuadratica = valores[2]
	m.MaxValor = valores[3]
}

func (m Matrix) Print() {
	fmt.Printf("Alto: %v \nAncho: %v \nCuadratica: %v \nMaxValor: %v \n", m.Alto, m.Ancho, m.Cuadratica, m.MaxValor)
}

func main() {
	mat := Matrix{}
	mat.Set(10.0, 10.0, 1.0, 5.0)
	mat.Print()
}
