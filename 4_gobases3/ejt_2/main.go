package main

import "fmt"

type Producto struct {
	name string
	price float64
	quantity int
}

type User struct {
	name string
	lastname string
	email string
	products []Producto
}

func newProduct(name string, price float64) Producto {
	return Producto{name, price, 0}
}

func addProduct(user *User, product Producto, quantity int) {
	user.products = append(user.products, Producto{product.name, product.price, quantity})
}

func cleanProducts(user *User) {
	user.products = []Producto{}
}

func main() {
	user := User{}
	user.name = "Juan"
	user.lastname = "Perez"
	user.email = "juan@mail.com"
	addProduct(&user, newProduct("Laptop", 1000), 1)
	addProduct(&user, newProduct("Smartphone", 500), 2)
	addProduct(&user, newProduct("Tablet", 500), 3)
	fmt.Println(user)
	cleanProducts(&user)
	fmt.Println(user)
}