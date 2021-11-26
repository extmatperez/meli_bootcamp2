package main

type usuario struct {
	Nombre   string
	Apellido string
	Correo   string
	producto []producto
}

type producto struct {
	Nombre   string
	Precio   int
	Cantidad int
}

func nuevoProducto(nombre string, precio int) producto {
	nuevProd := producto{Nombre: "cuchara", Precio: 3}
	return nuevProd
}

func agregarProducto(usuar usuario, produc producto, cantidad int) {
	//usuario =  producto{producto: producto}
}

func borrarProducto() {

}

func main() {

}
