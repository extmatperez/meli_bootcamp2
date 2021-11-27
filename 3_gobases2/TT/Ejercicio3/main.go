package main

import "fmt"

const (
	pequeño = "pequeño"
	mediano = "mediano"
	grande  = "grande"
)

type Producto interface {
	CalcularCosto() float64
}

type Ecommerce interface {
	Total() float64
	Agregar(Producto)
}

type producto struct {
	Tipo   string
	Nombre string
	Precio float64
}

func (p producto) CalcularCosto() (adicional float64) {
	switch p.Tipo {
	case pequeño:
		adicional = p.Precio
	case mediano:
		adicional = p.Precio + (p.Precio * 0.03)
	case grande:
		adicional = (p.Precio * 0.06) + 2500 + p.Precio
	}
	return
}

type Tienda struct {
	prod []Producto
}

func (t Tienda) Total() (total float64) {
	for _, p := range t.prod {
		total += p.CalcularCosto()
	}
	return
}

func (t *Tienda) Agregar(p Producto) {
	t.prod = append(t.prod, p)
}

func nuevoProducto(tipo, nombre string, precio float64) Producto {
	return producto{tipo, nombre, precio}
}

func nuevaTienda() Ecommerce {
	t := &Tienda{}
	return t
}

func main() {
	tienda1 := nuevaTienda()
	fmt.Println(tienda1)

	p1 := nuevoProducto(pequeño, "chocolate", 100)
	p2 := nuevoProducto(mediano, "impresora", 1000)
	p3 := nuevoProducto(grande, "nevera", 10000)

	fmt.Println(p1.CalcularCosto())
	fmt.Println(p2.CalcularCosto())
	fmt.Println(p3.CalcularCosto())

	tienda1.Agregar(p1)
	tienda1.Agregar(p2)
	tienda1.Agregar(p3)
	fmt.Println(tienda1)

	fmt.Println(tienda1.Total())
	//tienda1.Total()

}
