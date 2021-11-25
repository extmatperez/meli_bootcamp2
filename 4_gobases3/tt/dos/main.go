package main

import "fmt"

type Product struct {
	Name   string
	Price  float64
	Amount int
}

type User struct {
	Name     string
	LName    string
	Email    string
	Products []Product
}

func newProduct(name string, price float64) Product {
	return Product{Name: name, Price: price}
}

func addProduct(user *User, product *Product, amount int) {
	(*product).Amount = amount
	(*user).Products = append((*user).Products, *product)
}

func deleteProduct(user *User) {
	(*user).Products = []Product{}
}

func main() {
	newUser := User{Name: "Emilio", LName: "Ravenna", Email: "emiravenna@gmail.com", Products: []Product{}}
	fmt.Println("Inicial: ", newUser)

	product1 := newProduct("pepitas", 9.99)
	product2 := newProduct("anteojo", 4.25)
	product3 := newProduct("pala", 4.25)

	addProduct(&newUser, &product1, 5)
	addProduct(&newUser, &product2, 2)
	addProduct(&newUser, &product3, 8)
	fmt.Println("Productos agregados: ", newUser)

	deleteProduct(&newUser)
	fmt.Println("Productos eliminados: ", newUser)
}
