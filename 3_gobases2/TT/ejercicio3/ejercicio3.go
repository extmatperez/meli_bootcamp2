/*Varias tiendas de ecommerce necesitan realizar una funcionalidad en Go
para administrar productos y retornar el valor del precio total.
Las empresas tienen 3 tipos de productos:
Pequeño, Mediano y Grande. (Se espera que sean muchos más)
Existen costos adicionales por mantener el producto en el almacén de la tienda,
y costos de envío.

Sus costos adicionales son:
Pequeño: El costo del producto (sin costo adicional)
Mediano: El costo del producto + un 3% por mantenerlo en existencia en el almacén de la tienda.
Grande: El costo del producto + un 6%  por mantenimiento, y un costo adicional
por envío de $2500.

Requerimientos:
Crear una estructura “tienda” que guarde una lista de productos.
Crear una estructura “producto” que guarde el tipo de producto, nombre y precio
Crear una interface “Producto” que tenga el método “CalcularCosto”
Crear una interface “Ecommerce” que tenga los métodos “Total” y “Agregar”.
Se requiere una función “nuevoProducto” que reciba el tipo de producto, su nombre y
precio y devuelva un Producto.
Se requiere una función “nuevaTienda” que devuelva un Ecommerce.
Interface Producto:
El método “CalcularCosto” debe calcular el costo adicional según el tipo de producto.
Interface Ecommerce:
 - El método “Total” debe retornar el precio total en base al costo total de los productos
 y los adicionales si los hubiera.
 - El método “Agregar” debe recibir un producto y añadirlo a la lista de la tienda*/

package main

import "fmt"

func main() {
	tipoGrande := Grande{}
	tipoMediano := Mediano{}
	tipoChico := Chico{}
	prod1 := producto{tipoGrande, "Heladera", 16956.99}
	prod2 := producto{tipoMediano, "Computadora", 82365}
	prod3 := producto{tipoChico, "Celular", 67692}

	productosTienda := []producto{prod1, prod2, prod3}
	tienda1 := tienda{productosTienda}
	fmt.Printf("El costo total de los productos es de: $ %.2f\n", tienda1.Total())

	prod4 := producto{tipoGrande, "Lavarropas", 9785.99}

	tienda1.Agregar(prod4)
	fmt.Printf("El costo total de los productos es de: $ %.2f\n", tienda1.Total())
}

type tienda struct {
	Productos []producto
}

func (t *tienda) Agregar(nuevoProducto producto) {
	t.Productos = append(t.Productos, nuevoProducto)
}

func (t *tienda) Total() float64 {
	var total float64 = 0.0
	for _, prod := range t.Productos {
		fmt.Printf("El precio del producto %v es de: $%.2f\n", prod.Nombre, prod.CalcularCosto())
		total += prod.CalcularCosto()
	}
	return total
}

type producto struct {
	Tipo   Tipo
	Nombre string
	Precio float64
}

func (p *producto) CalcularCosto() float64 {
	return p.Tipo.CalcularAdicional()(p.Precio)
}

type Producto interface {
	CalcularCosto() float64
}

type Ecommerce interface {
	Total() float64
	Agregar(producto Producto)
}

type Tipo interface {
	CalcularAdicional() func(precio float64) float64
}

type Grande struct{}

func (n Grande) CalcularAdicional() func(precio float64) float64 {
	fun := func(precio float64) float64 {
		return precio + ((precio * 6) / 100) + 2500
	}
	return fun
}

type Mediano struct{}

func (n Mediano) CalcularAdicional() func(precio float64) float64 {
	fun := func(precio float64) float64 {
		return precio + ((precio * 3) / 100)
	}
	return fun
}

type Chico struct{}

func (n Chico) CalcularAdicional() func(precio float64) float64 {
	fun := func(precio float64) float64 {
		return precio
	}
	return fun
}
