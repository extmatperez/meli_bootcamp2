package main

import (
	"fmt"

)

type Producto interface {
	CalcularCosto() float64
}

type Ecommerce interface {
	Total() float64
	Agregar(p producto) tienda
}


type tienda struct {
	productos []producto

} 
type producto struct {
	Tipo,Nombre string
	Precio float64
} 

func (p *producto) NuevoProducto(tipo,nombre string, precio float64) producto{
	p.Nombre = nombre
	p.Precio= precio
	p.Tipo= tipo
	return *p
}

func (t tienda) NuevaTienda(productos ...producto) tienda{
	t.productos=productos
	return t;
}

func (p *producto) CalcularCosto() float64{
	switch p.Tipo{
	case "lacteo": 
	return p.Precio * 1.2
	default:
	return p.Precio
	}
}

func(t *tienda) Total() float64{
	 total := 0.0

	 for _, producto := range t.productos{
		total += producto.CalcularCosto()
	 }

	 return total
}

func(t *tienda) Agregar(p producto) tienda{
	t.productos = append(t.productos, p)

	return *t
}



func main() {

	producto1 := producto{Nombre: "leche",Precio: 100.0,Tipo: "lacteo"}
	producto2 := producto{Nombre: "queso",Precio: 100.0,Tipo: "lacteo"}
	producto3 := producto{Nombre: "galleta",Precio: 70.0,Tipo: "comestible"}

	tienda1 := tienda{}

	tienda1.Agregar(producto1)
	tienda1.Agregar(producto2)
	tienda1.Agregar(producto3)

	fmt.Println(tienda1)
	fmt.Println(tienda1.Total())

	producto4 := producto{Nombre: "yogut",Precio: 100.0,Tipo: "lacteo"}
	producto5 := producto{Nombre: "queso crema",Precio: 100.0,Tipo: "lacteo"}
	producto6 := producto{Nombre: "papas",Precio: 70.0,Tipo: "comestible"}

	tienda2 := tienda{}
	tienda2.Agregar(producto4)
	tienda2.Agregar(producto5)
	tienda2.Agregar(producto6)

	fmt.Println(tienda2)
	fmt.Println(tienda2.Total())

}

