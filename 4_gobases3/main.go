package main

import (
	"fmt"
	"math"
)

type Circulo struct {
	radio float64
}

//Agregar metodo de la estructura Circulo
func (c Circulo) area() float64 {
	return math.Pi * c.radio * c.radio
}

func (c Circulo) perim() float64 {
	return 2 * math.Pi * c.radio
}

func (c *Circulo) setRadio(r float64) {
	c.radio = r
}

func main() {
	circulo := Circulo{radio:10.0}
	fmt.Println(circulo.area())
	fmt.Println(circulo.perim())

	circulo.setRadio(20.0)
	fmt.Println(circulo.area())
	fmt.Println(circulo.perim())

}