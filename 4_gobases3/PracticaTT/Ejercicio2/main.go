package main

import "fmt"

/*Una importante empresa de ventas web necesita agregar una funcionalidad para agregar productos a los usuarios.
Para ello requieren que tanto los usuarios como los productos tengan la misma dirección de memoria en el main del
programa como en las funciones.
Se necesitan las estructuras:
Usuario: Nombre, Apellido, Correo, Productos (array de productos).
Producto: Nombre, precio, cantidad.
Se requieren las funciones:
Nuevo producto: recibe nombre y precio, y retorna un producto.
Agregar producto: recibe usuario, producto y cantidad, no retorna nada, agrega el producto al usuario.
Borrar productos: recibe un usuario, borra los productos del usuario.
*/

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
	prod := Producto{Nombre: nombre,
		Precio: precio,
	}

	return prod
}

func agregarProducto(usuario *Usuario, producto *Producto, cant int) {
	producto.Cantidad = cant
	usuario.Productos = append(usuario.Productos, *producto)
}

func borrarProductos(usuario *Usuario) {
	usuario.Productos = nil
}

func main() {
	prod1 := nuevoProducto("Alfajor", 50)

	usuario1 := Usuario{Nombre: "Facundo",
		Apellido: "Bouza",
		Correo:   "facubouza@gmail.com",
	}

	agregarProducto(&usuario1, &prod1, 15)
	fmt.Println("Usuario con un producto", usuario1)

	prod2 := nuevoProducto("Chicle", 30)
	agregarProducto(&usuario1, &prod2, 5)
	fmt.Println("Usuario con un producto", usuario1)

	borrarProductos(&usuario1)
	fmt.Println("Usuario con un producto", usuario1)

	prod3 := nuevoProducto("Caramelo", 20)
	agregarProducto(&usuario1, &prod3, 33)
	fmt.Println("Usuario con un producto", usuario1)
}
