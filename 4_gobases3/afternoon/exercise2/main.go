package main

import "fmt"

type Productos struct {
	Nombre   string `json:"nombre"`
	Precio   int    `json:"precio"`
	Cantidad int    `json:"cantidad"`
}
type Usuarios struct {
	Nombre    string `json:"nombre"`
	Apellido  string `json:"apellido"`
	Correo    string `json:"correo"`
	Productos []Productos
}

func nuevoProducto(nombre string, precio int) Productos {
	var producto Productos
	producto.Nombre = nombre
	producto.Precio = precio
	return producto
}

func agregarProducto(usuario *Usuarios, producto Productos, cantidad int) {
	producto.Cantidad = cantidad
	usuario.Productos = append(usuario.Productos, producto)
}
func borrarProducto(usuario *Usuarios) {
	usuario.Productos = nil
}

func main() {
	/*
		Una importante empresa de ventas web necesita agregar una funcionalidad para agregar productos a los usuarios. Para ello requieren que tanto los usuarios como los productos tengan la misma dirección de memoria en el main del programa como en las funciones.
		Se necesitan las estructuras:
		Usuario: Nombre, Apellido, Correo, Productos (array de productos).
		Producto: Nombre, precio, cantidad.
		Se requieren las funciones:
		Nuevo producto: recibe nombre y precio, y retorna un producto.
		Agregar producto: recibe usuario, producto y cantidad, no retorna nada, agrega el producto al usuario.
		Borrar productos: recibe un usuario, borra los productos del usuario.
	*/
	usuario1 := Usuarios{
		Nombre:   "Juan",
		Apellido: "Perez",
		Correo:   "asd",
	}
	usuario2 := Usuarios{"ariel", "romero", "rrr", nil}

	cuchara := nuevoProducto("cuchara", 10)
	tenedor := nuevoProducto("tenedor", 13)

	agregarProducto(&usuario1, cuchara, 2)

	fmt.Println(usuario1)

	agregarProducto(&usuario1, tenedor, 4)

	fmt.Println(usuario1)

	borrarProducto(&usuario1)

	fmt.Println(usuario1)
	fmt.Println(usuario2)
}
