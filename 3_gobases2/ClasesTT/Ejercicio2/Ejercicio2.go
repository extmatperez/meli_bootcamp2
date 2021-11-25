package main

import (
	"encoding/json"
	"fmt"
)

type Matrix struct {
	Matriz     float64
	Width      float64
	Height     float64
	Cuadratica bool
	MaxValue   float64
}

func (matrix *Matrix) Set(Matriz, Width, Height, MaxValue float64, Cuadratica bool) {
	matrix.Matriz = Matriz
	matrix.Width = Width
	matrix.Height = Height
	matrix.Cuadratica = Cuadratica
	matrix.MaxValue = MaxValue
}

func (matrix Matrix) Print() {
	miJson, _ := json.Marshal(matrix)
	fmt.Println(string(miJson))
}
func main() {
	m := Matrix{}
	m.Set(3.0, 4.0, 4.0, 4.0, true)
	m.Print()
}
