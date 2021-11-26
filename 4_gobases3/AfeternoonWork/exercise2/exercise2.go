/*
Una importante empresa de ventas web necesita agregar una funcionalidad para agregar productos a los usuarios.
Para ello requieren que tanto los usuarios como los productos tengan la misma dirección de memoria en el main
del programa como en las funciones.

Se necesitan las estructuras:

Usuario: Nombre, Apellido, Correo, Productos (array de productos).
Producto: Nombre, precio, cantidad.

Se requieren las funciones:

Nuevo producto: recibe nombre y precio, y retorna un producto.
Agregar producto: recibe usuario, producto y cantidad, no retorna nada, agrega el producto al usuario.
Borrar productos: recibe un usuario, borra los productos del usuario.

*/

package main

import "fmt"

type Product struct {
	Name     string
	Price    float64
	Quantity int
}

type UsersNew struct {
	Name     string
	LastName string
	Email    string
	Product  []Product
}

func (user UsersNew) getProduct() []Product {
	return user.Product
}

func newProduct(newName string, newPrice float64) Product {
	prod := Product{}
	prod.Name = newName
	prod.Price = newPrice
	return prod
}

func addProduct(user *UsersNew, product Product, quantity int) {
	product.Quantity = quantity
	listProd := user.getProduct()
	listProd = append(listProd, product)
	user.Product = listProd
}

func main() {

	producto1 := newProduct("Mouse", 70.99)
	producto2 := newProduct("Monitor", 150.11)
	producto3 := newProduct("MAC", 999.99)

	usuario1 := UsersNew{Name: "Jose", LastName: "Rios", Email: "joserios@mercadolibre.cl"}
	addProduct(&usuario1, producto1, 5)
	addProduct(&usuario1, producto2, 3)
	addProduct(&usuario1, producto3, 1)
	fmt.Println(usuario1)
}
