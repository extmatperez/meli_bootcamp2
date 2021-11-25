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

type Usuario struct {
	nombre    string
	apellido  string
	correo    string
	productos []Producto
}

type Producto struct {
	nombre   string
	precio   float64
	cantidad int
}

func nuevoProducto(nombre string, precio float64) Producto {
	return Producto{nombre, precio, 0}
}

func agregarProducto(usuario *Usuario, producto *Producto, cantidad int) {
	producto.cantidad = cantidad

	usuario.productos = append(usuario.productos, *producto)
}

func borrarProducto(usuario *Usuario) {
	usuario.productos = nil
}

func main() {

	u := Usuario{"Pepe", "Flores", "pepitoflores@gmail.com", nil}

	fmt.Println(u)

	p := nuevoProducto("Coca Cola", 10)

	agregarProducto(&u, &p, 5)

	p = nuevoProducto("Pepsi", 15)

	agregarProducto(&u, &p, 15)

	fmt.Println(u)

	borrarProducto(&u)

	fmt.Println(u)
}
