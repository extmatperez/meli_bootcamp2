package main

import "fmt"

func main() {
	user := User{
		Name:     "Matias",
		Lastname: "Ziliotto",
		Email:    "matias.ziliotto@mercadolibre.com",
	}

	product1 := NewProduct("Carne", 650.00)
	product2 := NewProduct("Coca-Cola", 130.00)
	product3 := NewProduct("Servilletas", 115.00)

	fmt.Println(user)
	fmt.Println("")

	AddProduct(&user, &product1, 10)
	AddProduct(&user, &product2, 20)
	AddProduct(&user, &product3, 5)

	fmt.Println(user)
	fmt.Println("")

	DeleteProducts(&user)

	fmt.Println(user)
}

type User struct {
	Name     string
	Lastname string
	Email    string
	Products []Product
}

type Product struct {
	Name   string
	Price  float64
	Amount int
}

func NewProduct(name string, price float64) Product {
	return Product{
		Name:  name,
		Price: price,
	}
}

func AddProduct(user *User, product *Product, amount int) {
	product.Amount = amount
	user.Products = append(user.Products, *product)
}

func DeleteProducts(user *User) {
	user.Products = []Product{}
}
