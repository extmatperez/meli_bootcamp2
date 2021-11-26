/*
? Ejercicio 2 - Ecommerce
Una importante empresa de ventas web necesita agregar una funcionalidad para
agregar productos a los usuarios. Para ello requieren que tanto los usuarios
como los productos tengan la misma dirección de memoria en el main del programa
como en las funciones.
Se necesitan las estructuras:
Usuario: Nombre, Apellido, Correo, Productos (array de productos).
Producto: Nombre, precio, cantidad.
Se requieren las funciones:
Nuevo producto: recibe nombre y precio, y retorna un producto.
Agregar producto: recibe usuario, producto y cantidad, no retorna nada, agrega
el producto al usuario.
Borrar productos: recibe un usuario, borra los productos del usuario.

*/

package main

import "fmt"

type Usuario struct {
	Nombre    string
	Apellido  string
	Correo    string
	Productos []Producto
}

type Producto struct {
	Nombre   string
	Precio   float64
	Cantidad int
}

func nuevoProducto(nombre string, precio float64) Producto {
	return Producto{Nombre: nombre, Precio: precio}
}

func agregarProducto(usuario *Usuario, producto *Producto, cantidad int) {
	(*producto).Cantidad = cantidad
	(*usuario).Productos = append((*usuario).Productos, *producto)

}

func borrarProducto(usuario *Usuario) {
	(*usuario).Productos = []Producto{}
}

func main() {

	usuario1 := Usuario{
		Nombre:   "Nico",
		Apellido: "Arguello",
		Correo:   "nico@gmail.com",
	}
	prod1 := nuevoProducto("coca", 150.50)
	prod2 := nuevoProducto("agua", 100.50)
	prod3 := nuevoProducto("cerveza", 120)

	agregarProducto(&usuario1, &prod1, 1)
	agregarProducto(&usuario1, &prod2, 2)
	agregarProducto(&usuario1, &prod3, 3)

	fmt.Printf("%+v\n", usuario1)
	fmt.Printf("%+v\n", prod1)

	borrarProducto(&usuario1)
	fmt.Printf("%+v\n", usuario1)

}
