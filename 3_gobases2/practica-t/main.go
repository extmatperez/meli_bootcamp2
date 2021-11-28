/* package main

import (
	"fmt"
	"math"
)

type Circulo struct {
	radio float64
}

func (c Circulo) area() float64 {
	return math.Pi * c.radio * c.radio
}

func (c Circulo) perim() float64 {
	return 2 * math.Pi * c.radio
}

func (c *Circulo) setRadio(r float64) {
	r = c.radio
}

func main() {

	c := Circulo{radio: 5}
	fmt.Println(c)
	fmt.Printf("%.2f\n", c.area())
	fmt.Printf("%.2f\n", c.perim())
	c.setRadio(10)
	fmt.Printf("%.2f\n", c.area())
	fmt.Printf("%.2f\n", c.perim())
}
*/

/* package main

import "fmt"

type Vehiculo struct {
	km     float64
	tiempo float64
}

func (v Vehiculo) detalle() {

	fmt.Printf("km:\t%f\ntiempo:\t%f\n", v.km, v.tiempo)
}

type Auto struct {
	v Vehiculo
}

func (a *Auto) Correr(minutos int) {
	a.v.tiempo = float64(minutos) / 60
	a.v.km = a.v.tiempo * 100
}

func (a *Auto) Detalle() {
	fmt.Println("\nV:\tAuto")
	a.v.detalle()
}

type Moto struct {
	v Vehiculo
}

func (m *Moto) Correr(minutos int) {
	m.v.tiempo = float64(minutos) / 60
	m.v.km = m.v.tiempo * 100
}

func (m *Moto) Detalle() {
	fmt.Println("\nV:\tMoto")
	m.v.detalle()
}

func main() {

	auto := Auto{}
	auto.Correr(360)
	auto.Detalle()

	moto := Moto{}
	moto.Correr(400)
	moto.Detalle()

}
*/

package main

import (
	"fmt"
	"math"
)

type geometry interface {
	area() float64
	perim() float64
}

type circle struct {
	radius float64
}

func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func (c circle) perim() float64 {
	return 2 * math.Pi * c.radius
}

type rect struct {
	width, height float64
}

func (r rect) area() float64 {
	return r.width * r.height
}

func (r rect) perim() float64 {
	return 2*r.width + 2*r.height
}

func details(g geometry) {
	fmt.Println(g)
	fmt.Println(g.area())
	fmt.Println(g.perim())
}

func main() {
	c := circle{radius: 5}
	r := rect{width: 2, height: 6}
	details(c)
	details(r)
}
