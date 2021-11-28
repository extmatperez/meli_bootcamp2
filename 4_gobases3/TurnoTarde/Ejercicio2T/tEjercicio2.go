package main

import (
	"fmt"

)
type Producto struct {
	Nombre string 
	Precio float64
	Cantidad int
} 
type Usuario struct {
	Prodcutos [] Producto
	Nombre,Apellido,Correo string 
} 



func (m *Producto) NewProduct(nombre string, precio float64) Producto{
	m.Nombre=nombre
	m.Precio=precio
	return *m
}
func (m *Usuario) AddProducto(prod Producto,cantidad int) {
	prod.Cantidad=cantidad
	m.Prodcutos = append(m.Prodcutos,prod)
}

func (m *Usuario) DeleteProducto(u Usuario) {
	m.Prodcutos = nil
}

func main() {

	prod1 := Producto{Nombre: "Carne", Precio: 55.8 }
	prod2 := Producto{Nombre: "Leche", Precio: 80.8 }
	prod3 := Producto{Nombre: "Papa", Precio: 15.8 }

	usuario := Usuario{Nombre: "Francisco",Apellido: "Palacio",Correo: "123456@yahoo.com"}
	usuario2 := Usuario{Nombre: "Carlos",Apellido: "Palaciossss",Correo: "123456@yahoo.com"}
	usuario.AddProducto(prod1,5)
	usuario.AddProducto(prod2,10)

	fmt.Println(usuario)

	usuario.DeleteProducto(usuario)

	fmt.Println(usuario)


	usuario2.AddProducto(prod1,5)
	usuario2.AddProducto(prod2,80)
	usuario2.AddProducto(prod3,50)

	fmt.Println(usuario2)

	usuario2.DeleteProducto(usuario2)

	fmt.Println(usuario2)

}

