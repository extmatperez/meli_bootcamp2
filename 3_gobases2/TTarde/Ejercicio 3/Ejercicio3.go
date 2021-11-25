package main

import "fmt"

const (
	Pequeño = "Pequeño"
	Mediano = "Mediano"
	Grande  = "Grande"
)

type producto struct {
	tipo      string
	nombre    string
	precio    float64
	adicional float64
}

type Producto interface {
	CalcularCosto()
}

func (p *producto) CalcularCosto() {
	switch p.tipo {
	case Mediano:
		p.adicional = p.precio / 100 * 3
	case Grande:
		p.adicional = (p.precio / 100 * 6) + 2500
	default:
		p.adicional = 0
	}
}

type tienda struct {
	lista []producto
}

type Ecommerce interface {
	Total() float64
	Agregar(p producto)
}

func (t tienda) Total() float64 {
	total := 0.0

	for i := 0; i < len(t.lista); i++ {
		total += t.lista[i].precio + (t.lista[i].adicional)
	}
	return total
}

func (t *tienda) Agregar(p producto) {
	t.lista = append(t.lista, p)
}

func nuevoProducto(tipo string, nombre string, precio float64) Producto {
	return &producto{tipo, nombre, precio, 0}
}

func nuevaTienda(lista []producto) Ecommerce {
	return &tienda{lista}
}

func main() {
	p := producto{nombre: "Peine", tipo: "Pequeño", precio: 244}
	p1 := producto{nombre: "Puerta", tipo: "Grande", precio: 690}
	list := []producto{p, p1}
	t := tienda{list}
	fmt.Println(t.Total())

	c := nuevoProducto("Moto", "Grande", 54410)
	c.CalcularCosto()
	t.lista[0].CalcularCosto()
	t.lista[1].CalcularCosto()

	t1 := nuevaTienda(list)
	fmt.Println(t1.Total())

	fmt.Println(t.Total())

}
