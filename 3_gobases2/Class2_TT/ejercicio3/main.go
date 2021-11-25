package main

const (
	pequeño = "pequeño"
	mediano = "mediano"
	grande  = "grande"
)

func main() {
	//productosTienda := 
	tienda1 := nuevaTienda(productosTienda, "tienda1")
}

type tienda struct {
	productos []producto
}

type producto struct {
	nombre       string
	precio       float64
	tipoProducto string
}

type Producto interface {
	calcularCosto() float64
}

type Ecommerce interface {
	total() float64
	agregar(producto producto) float64
}

func (p *producto) calcularCosto() float64 {
	return 0
}

func (t *tienda) total() float64 {
	return 0
}

func (t *tienda) agregar(nuevoProduct producto) float64 {
	return 0
}

func nuevoProduct() producto {
	return producto{}
}

func nuevaTienda(productosTienda []producto, tipo string) Ecommerce {
	switch tipo {
	case "tienda1":
		return &tienda{productosTienda,
		}
}
