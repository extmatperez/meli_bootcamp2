/*Una importante empresa de ventas web necesita agregar una funcionalidad para agregar productos a los usuarios. Para ello
requieren que tanto los usuarios como los productos tengan la misma dirección de memoria en el main del programa como en las
funciones.
Se necesitan las estructuras:
Usuario: Nombre, Apellido, Correo, Productos (array de productos).
Producto: Nombre, precio, cantidad.
Se requieren las funciones:
Nuevo producto: recibe nombre y precio, y retorna un producto.
Agregar producto: recibe usuario, producto y cantidad, no retorna nada, agrega el producto al usuario.
Borrar productos: recibe un usuario, borra los productos del usuario.
*/

package main

var user *string
var v string

type User struct {
	Nombre, Apellido, Correo, Productos string
}

type Product struct {
	Nombre   string
	Precio   float64
	Cantidad int
}

func nuevoProducto(nombre string, precio float64) string {

}

func agregarProducto(user, producto string, cantidad int) {

}

func borrarProductos(user string) {

}

func main() {
	user = &v

}
