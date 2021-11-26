/*
Una empresa nacional se encarga de realizar venta de productos, servicios y mantenimiento.
Para ello requieren realizar un programa que se encargue de calcular el precio total de Productos, Servicios y Mantenimientos.
Debido a la fuerte demanda y para optimizar la velocidad requieren que el cálculo de la sumatoria se realice en paralelo mediante 3 go routines.

Se requieren 3 estructuras:
Productos: nombre, precio, cantidad.
Servicios: nombre, precio, minutos trabajados.
Mantenimiento: nombre, precio.

Se requieren 3 funciones:
Sumar Productos: recibe un array de producto y devuelve el precio total (precio * cantidad).
Sumar Servicios: recibe un array de servicio y devuelve el precio total (precio * media hora trabajada, si no llega a trabajar 30 minutos se le cobra como si hubiese trabajado media hora).
Sumar Mantenimiento: recibe un array de mantenimiento y devuelve el precio total.

Los 3 se deben ejecutar concurrentemente y al final se debe mostrar por pantalla el monto final (sumando el total de los 3).


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

func deleteAllProduct(user *UsersNew) {
	productEmpty := Product{}
	var listEmpty []Product
	listEmpty = append(listEmpty, productEmpty)
	user.Product = listEmpty
}

func main() {

	producto1 := newProduct("Mouse", 70.99)
	producto2 := newProduct("Monitor", 150.11)
	producto3 := newProduct("MAC", 999.99)

	usuario1 := UsersNew{Name: "Jose", LastName: "Rios", Email: "joserios@mercadolibre.cl"}
	usuario2 := UsersNew{Name: "Paula", LastName: "Isabel", Email: "paulaisabel@mercadolibre.cl"}
	addProduct(&usuario1, producto1, 5)
	addProduct(&usuario1, producto2, 3)
	addProduct(&usuario2, producto1, 5)
	addProduct(&usuario2, producto2, 10)
	addProduct(&usuario1, producto3, 1)
	fmt.Println(usuario1)

	deleteAllProduct(&usuario1)
	fmt.Println(usuario2)
	fmt.Println(usuario1)

}
