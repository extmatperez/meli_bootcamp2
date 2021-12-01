package main

import "fmt"

// Ejercicio 3 - Productos
// Varias tiendas de ecommerce necesitan realizar una funcionalidad en Go para administrar productos y retornar el valor del precio total.
// Las empresas tienen 3 tipos de productos:
// Pequeño, Mediano y Grande. (Se espera que sean muchos más)
// Existen costos adicionales por mantener el producto en el almacén de la tienda, y costos de envío.

// Sus costos adicionales son:
// Pequeño: El costo del producto (sin costo adicional)
// Mediano: El costo del producto + un 3% por mantenerlo en existencia en el almacén de la tienda.
// Grande: El costo del producto + un 6%  por mantenimiento, y un costo adicional  por envío de $2500.

// type ProdSmall struct{}

// type ProdMedium struct{}

// type ProdBig struct{}

// Crear una estructura “tienda” que guarde una lista de productos.

type Shop struct {
	ProductsList []Product
}

// Crear una estructura “producto” que guarde el tipo de producto, nombre y precio

type Product struct {
	ProductType string
	Name        string
	Price       float64
}

// Se requiere una función “nuevoProducto” que reciba el tipo de producto, su nombre y precio y devuelva un Producto.

func newProduct(ProductType, Name string, Price float64) Product {
	return Product{
		ProductType,
		Name,
		Price,
	}

}

// Crear una interface “Producto” que tenga el método “CalcularCosto”
// Interface Producto:
// El método “CalcularCosto” debe calcular el costo adicional según el tipo de producto.

type ProductInterface interface {
	calculateCost()
}

func calculateCost(p Product) float64 {
	switch p.ProductType {
	case "small":
		return p.Price
	case "medium":
		return p.Price + p.Price*0.03
	case "big":
		return (p.Price + p.Price*0.06) + 2500
	}
	return 0
}

// Crear una interface “Ecommerce” que tenga los métodos “Total” y “Agregar”.
// Interface Ecommerce:
//  - El método “Total” debe retornar el precio total en base al costo total de los productos y los adicionales si los hubiera.
//  - El método “Agregar” debe recibir un producto y añadirlo a la lista de la tienda

// Se requiere una función “nuevaTienda” que devuelva un Ecommerce.
type Ecommerce interface {
	Total() float64
	Add() Shop
}

// func newShop(e *Ecommerce) Ecommece {

// }

// func Add(s *Shop) (product Product) {
// 	s.ProductsList = append(s.ProductsList, Product{})

// }

func main() {

	newprod := newProduct("big", "tv", 34000)
	fmt.Println(newprod)

	fmt.Println("\nthe cost is")

	cost := calculateCost(newprod)
	fmt.Println(cost)

}
