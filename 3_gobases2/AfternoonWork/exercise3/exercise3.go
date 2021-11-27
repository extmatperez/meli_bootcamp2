/*
Varias tiendas de ecommerce necesitan realizar una funcionalidad en Go para administrar productos y retornar el valor del precio total.
Las empresas tienen 3 tipos de productos:
Pequeño, Mediano y Grande. (Se espera que sean muchos más)
Existen costos adicionales por mantener el producto en el almacén de la tienda, y costos de envío.

Sus costos adicionales son:
Pequeño: El costo del producto (sin costo adicional)
Mediano: El costo del producto + un 3% por mantenerlo en existencia en el almacén de la tienda.
Grande: El costo del producto + un 6%  por mantenimiento, y un costo adicional  por envío de $2500.

Requerimientos:
ok Crear una estructura “tienda” que guarde una lista de productos.
ok Crear una estructura “producto” que guarde el tipo de producto, nombre y precio
ok Crear una interface “Producto” que tenga el método “CalcularCosto”
ok Crear una interface “Ecommerce” que tenga los métodos “Total” y “Agregar”.
ok Se requiere una función “nuevoProducto” que reciba el tipo de producto, su nombre y precio y devuelva un Producto.
Se requiere una función “nuevaTienda” que devuelva un Ecommerce.
Interface Producto:
El método “CalcularCosto” debe calcular el costo adicional según el tipo de producto.
Interface Ecommerce:
 - El método “Total” debe retornar el precio total en base al costo total de los productos y los adicionales si los hubiera.
 - El método “Agregar” debe recibir un producto y añadirlo a la lista de la tienda


*/
package main

import (
	"errors"
	"fmt"
)

const (
	LG = "lg"
	MD = "md"
	SM = "sm"
)

type Store struct {
	Products []ProductInterface
}

func (s Store) Total() float64 {
	var total float64
	for i := 0; i < len(s.Products); i++ {
		total += s.Products[i].CalculateCost()
	}
	return total
}

func (s *Store) AddProduct(productInterfaces ...ProductInterface) {
	for i := 0; i < len(productInterfaces); i++ {
		s.Products = append(s.Products, productInterfaces[i])
	}
}

type Product struct {
	Type  string
	Name  string
	Price float64
}

//Metodo de producto heredado de Producto
func (producto Product) CalculateCost() float64 {
	var cost float64 = producto.Price
	if producto.Type == MD {
		return cost * 1.03
	} else if producto.Type == LG {
		return (cost * 1.06) + 2500
	}
	return cost
}

type ProductInterface interface {
	CalculateCost() float64
}

type EcommerceInterface interface {
	Total() float64
	AddProduct(productInterfaces ...ProductInterface)
}

func newProduct(typeProd, nameProd string, priceProd float64) (ProductInterface, error) {
	switch typeProd {
	case LG:
		return Product{
			Type:  typeProd,
			Name:  nameProd,
			Price: priceProd,
		}, nil
	case MD:
		return Product{Type: typeProd, Name: nameProd, Price: priceProd}, nil
	case SM:
		return Product{Type: typeProd, Name: nameProd, Price: priceProd}, nil
	default:
		errorMsg := "el producto '" + typeProd + "' no existe"
		return Product{}, errors.New(errorMsg)
	}
}

func newStore() EcommerceInterface {
	return &Store{}
}

func main() {
	store := newStore()

	producto1, err := newProduct(LG, "Televisor", 920.90)

	if err != nil {
		fmt.Println(err)
	} else {
		store.AddProduct(producto1)

		fmt.Println(store.Total())
	}

}
