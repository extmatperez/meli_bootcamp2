package main

import (
	"fmt"
)

type Products struct {
	Name 		string
	Price		int
	Quantity	int
}

type User struct {
	Name 		string
	LastName 	string
	Email		string
	Products	[]Products
}

func main() {
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

	var pList []Products
	// set products
	p := newProduct("cucharas", 500)
	p2 := newProduct("tenedores", 500)
	pList = append(pList, p)
	//set User
	u := User{"juan", "soza", "JuanZosa@mail.com", nil}

	addProduct(&u, &p, 5)
	addProduct(&u, &p2, 4)
	fmt.Println("El usuario: ", u)

	addProduct(&u, &p, 5)
	fmt.Println("El usuario Modificado: ", u)

	deleteProducts(&u)
	fmt.Println("El usuario sin productos", u)
}
func newProduct(name string, price int) Products {
	nProduct := Products{name, price, 0}
	return nProduct
}
func addProduct(user *User, product *Products, quantity int)  {
	checker := false
	for i, value := range user.Products{
		if value.Name == product.Name {
			user.Products[i].Quantity = user.Products[i].Quantity + quantity
			user.Products[i].Price = product.Price * user.Products[i].Quantity
			checker = true
		}
	}
	if checker == false {
		product.Quantity = quantity
		user.Products = append(user.Products, *product)
		for i, _ := range user.Products{
				user.Products[i].Price = product.Price * user.Products[i].Quantity
		}
	}
}
func deleteProducts(user *User) {
	user.Products = nil
}