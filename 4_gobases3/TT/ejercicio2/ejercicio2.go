package main

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
import (
	"fmt"
)

type Usuario struct {
	Nombre    string
	Apellido  string
	Correo    string
	Productos []Product
}
type Product struct {
	Nombre   string
	Precio   float64
	Cantidad int
}

func newProduct(product string, price float64) *Product {
	var temp Product
	temp.Nombre = product
	temp.Precio = price
	temp.Cantidad = 0
	return &temp
}

func addProduct(usuario *Usuario, producto *Product, cant int) {
	(*producto).Cantidad += cant
	(*usuario).Productos = append((*usuario).Productos, *producto)
}

func emptyProduct(usuario *Usuario) {
	var empty = []Product{}
	(*usuario).Productos = empty
}

func main() {
	fmt.Println("Ejercicio2")
	var user Usuario
	var prod1 Product
	var prod2 *Product
	prod1.Nombre = "Bananas"
	prod1.Precio = 215
	prod1.Cantidad = 12
	prod2 = newProduct("Manzanas", 155)

	user.Nombre = "Juan"
	user.Apellido = "Perez"
	user.Correo = "as@hotmail.com"
	user.Productos = append(user.Productos, prod1)

	fmt.Println("Usuario Inicial: ", user)
	addProduct(&user, prod2, 6)
	fmt.Println("Usuario con productos: ", user)
	emptyProduct(&user)
	fmt.Println("Usuario final: ", user)
}
