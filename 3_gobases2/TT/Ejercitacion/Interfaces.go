package main

import (
	"fmt"
	"math"
)

type geometry interface {
	area() float64
	perimeter() float64
}

// Para rectangulos.
type rectangle struct {
	width  float64
	height float64
}

func (r rectangle) area() float64 {
	return r.width * r.height
}

func (r rectangle) perimeter() float64 {
	return (2 * r.width) + (2 * r.height)
}

// Para circulos.
type circle struct {
	radius float64
}

func (c circle) area() float64 {
	return math.Pi * math.Pow(2, c.radius)
}

func (c circle) perimeter() float64 {
	return 2 * math.Pi * c.radius
}

// Para cuadrados.
type square struct {
	arist float64
}

func (s square) area() float64 {
	return math.Pow(2, s.arist)
}

func (s square) perimeter() float64 {
	return 4 * s.arist
}

// Para triangulos.
type triangle struct {
	base   float64
	height float64
}

func (t triangle) area() float64 {
	return (t.base * t.height) / 2
}

// Siempre y cuando sea equilatero el triangulo.
func (t triangle) perimeter() float64 {
	return 3 * t.base
}

// Para obtener los detalles con la interfaz geometry.
func getDetails(g geometry) {
	fmt.Println(g)
	fmt.Println(g.area())
	fmt.Println(g.perimeter())
}

func main() {
	r := rectangle{height: 5, width: 10}
	c := circle{radius: 4}
	getDetails(r)
	getDetails(c)
}
