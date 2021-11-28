package main

import "fmt"

//? CLASE PADRE
type Vehiculo struct {
	km     float64
	tiempo float64
}

// ! METODO CLASE PADRE
func (v Vehiculo) detalle() {
	fmt.Printf("km: \t %8.2f \ntiempo: %8.2f \n", v.km, v.tiempo)
}

//? CLASE HIJA
type Auto struct {
	v Vehiculo // * campo de TIPO VEHICULO
}
type Moto struct {
	v Vehiculo // * campo de TIPO VEHICULO
}

// ! METODO AUTO para 100 km/h
func (a *Auto) Correr(min int) {
	a.v.tiempo = float64(min) / 60
	a.v.km = a.v.tiempo * 100
}

// ! METODO Detalle que llama al metodo de la clase PADRE
func (a *Auto) Detalle() {
	fmt.Println("\n V: \t Auto")
	a.v.detalle()
}

// ! METODO MOTO para 80 km/h
func (m *Moto) Correr(min int) {
	m.v.tiempo = float64(min) / 60
	m.v.km = m.v.tiempo * 80
}

// ! METODO Detalle que llama al metodo de la clase PADRE
func (m *Moto) Detalle() {
	fmt.Println("\n V: \t Moto")
	m.v.detalle()
}

func main() {
	auto := Auto{}
	auto.Correr(360)
	auto.Detalle()

	moto := Moto{}
	moto.Correr(360)
	moto.Detalle()
	// fmt.Println(auto)
}
