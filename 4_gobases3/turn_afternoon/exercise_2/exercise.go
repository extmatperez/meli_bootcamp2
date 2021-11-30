package main

import "fmt"

func main() {
	user := User{"Matilda", "Rolan", nil, "madilta@example.com"}

	prod1 := createProduct("apple", 200.00)
	prod2 := createProduct("pear", 300.00)
	prod3 := createProduct("peach", 500.00)

	products := []Product{prod1, prod2, prod3}

	user.products = products

	fmt.Println("New User:", user)

	user = User{"Dig", "Davila", nil, "dig@example.com"}
	prod4 := createProduct("nuts", 120.00)

	setProductToUser(&user, &prod4, 1000)

	fmt.Println("Setter User:", user)

	deleteProductToUser(&user)

	fmt.Println("Delete product to User:", user)

}

type User struct {
	name     string
	lastName string
	products []Product
	email    string
}

type Product struct {
	name     string
	price    float64
	quantity int
}

func createProduct(name string, price float64) Product {
	return Product{name, price, 0}

}

func setProductToUser(user *User, prod *Product, quantity int) {
	prod.quantity = quantity
	prods := []Product{*prod}
	user.products = append(user.products, prods[0])
}

func deleteProductToUser(user *User) {
	user.products = nil
}
