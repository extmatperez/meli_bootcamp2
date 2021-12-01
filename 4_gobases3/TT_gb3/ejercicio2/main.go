package main

import (
	"fmt"
)

// Ejercicio 2 - Ecommerce
// Una importante empresa de ventas web necesita agregar una funcionalidad para agregar productos a los usuarios. Para ello requieren que tanto los usuarios como los productos tengan la misma dirección de memoria en el main del programa como en las funciones.
// Se necesitan las estructuras:
// Usuario: Nombre, Apellido, Correo, Productos (array de productos).
// Producto: Nombre, precio, cantidad.
// Se requieren las funciones:
// Nuevo producto: recibe nombre y precio, y retorna un producto.
// Agregar producto: recibe usuario, producto y cantidad, no retorna nada, agrega el producto al usuario.
// Borrar productos: recibe un usuario, borra los productos del usuario.

type User struct {
	Name, Surname, Mail string
	Products            []Product
}

type Product struct {
	Name     string
	Price    float64
	Quantity float64
}

// func newProduct(name string, price float64, p *Product) {
// 	(*p).Name = name
// 	(*p).Price = price

// 	return p
// }

func addProduct(p *Product, quantity float64, u *User) {
	(*p).Quantity = quantity
	(*u).Products = append((*u).Products, *p)
}

func newProduct(name string, price float64, p *Product) Product {
	(*p) = Product{Name: name, Price: price}

	return *p
}

func main() {

	prod1 := newProduct("monitor", 33000, &Product{})
	fmt.Println(prod1)

	user1 := User{}
	addProduct(&prod1, 3, &User{})
	fmt.Printf("\nafter added: %v", user1)

}
