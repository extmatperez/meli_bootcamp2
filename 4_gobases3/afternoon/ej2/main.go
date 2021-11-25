package main

import (
	"fmt"
)

type Product struct {
	name  string
	price float64
	count int
}

type User struct {
	name     string
	lastName string
	email    string
	products []Product
}

func newProduct(name string, price float64) *Product {
	return &Product{name: name, price: price, count: 0}
}

func addProduct(product *Product, user *User, count int) {
	product.count = count
	user.products = append(user.products, *product)
}

func removeProducts(user *User) {
	user.products = user.products[:0]
	fmt.Printf("%v\n", user)
	fmt.Printf("%p\n", user)
}

func main() {
	var u = new(User)
	fmt.Printf("%v\n", u)
	removeProducts(u)

}
