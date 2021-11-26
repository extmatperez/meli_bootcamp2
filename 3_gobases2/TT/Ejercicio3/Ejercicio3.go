package main

import (
	"fmt"
)

const (
	prodPequeno = "PEQUENO"
	prodMediano = "MEDIANO"
	prodGrande  = "GRANDE"
)

// Para tienda y gestion de la interfaz de ecommerce.
type tienda struct {
	Productos []producto `json:"productos"`
}

func (t tienda) total() float64 {
	total := 0.0
	for _, element := range t.Productos {
		total += element.calcular_costo()
	}
	return total
}

func (t *tienda) agregar(p producto) {
	t.Productos = append(t.Productos, p)
}

type producto struct {
	Tipo_Producto string  `json:"tipo_producto"`
	Nombre        string  `json:"nombre"`
	Precio        float64 `json:"precio"`
}

func (p producto) calcular_costo() float64 {
	var tipos_producto = map[string]float64{
		"PEQUENO": p.Precio,
		"MEDIANO": p.Precio * 1.03,
		"GRANDE":  (p.Precio * 1.06) + 2500.00,
	}
	return tipos_producto[p.Tipo_Producto]
}

type Producto interface {
	calcular_costo() float64
}

type ecommerce interface {
	total() float64
	agregar(p producto)
}

func nuevo_producto(prodType string, nombre string, precio float64) producto {
	return producto{Tipo_Producto: prodType, Nombre: nombre, Precio: precio}
}

func nueva_tienda() ecommerce {
	var products []producto = make([]producto, 0)
	var tda = tienda{products}
	var eCom ecommerce = &tda
	return eCom
}

func main() {
	p2 := nuevo_producto(prodPequeno, "Tuerca", 0.50)
	fmt.Println("Costo producto pequeño 2: ", p2.calcular_costo())

	m := nuevo_producto(prodMediano, "Martillo", 5.50)
	fmt.Println("Costo producto mediano: ", m.calcular_costo())

	g := nuevo_producto(prodGrande, "Amoladora", 55.60)
	fmt.Println("Costo producto grande: ", g.calcular_costo())

	var tiendita = nueva_tienda()
	tiendita.agregar(m)
	tiendita.agregar(g)
	tiendita.agregar(p2)

	fmt.Println("Total de productos de la tienda: ", tiendita.total())

}
