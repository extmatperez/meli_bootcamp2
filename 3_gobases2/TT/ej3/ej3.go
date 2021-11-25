package ej3

import "fmt"

type Product struct {
	Type  string
	Name  string
	Price float64
}

type Shop struct {
	Products []Product
}

type IProduct interface {
	CalculateCost()
}

type IEccomerce interface {
	Total()
	Add()
}

func NewProduct(productType, name string, price float64) Product {
	newProduct := Product{productType, name, price}

	return newProduct
}

func NewShop(products ...Product) Shop {
	newShop := Shop{products}

	return newShop
}

func (product *Product) CalculateCost() float64 {
	additionalCosts := map[string]float64{
		"sm": product.Price,
		"md": product.Price + (product.Price * 0.03),
		"lg": product.Price + (product.Price * 0.06) + 2500,
	}

	return additionalCosts[product.Type]
}

func (shop Shop) Total() float64 {
	result := 0.0

	for _, el := range shop.Products {
		result += el.Price
	}

	return result
}

func (shop *Shop) Add(product Product) []Product {
	shop.Products = append(shop.Products, product)

	return shop.Products
}

func Ej3() {
	product1 := NewProduct("sm", "Pen", 1.5)
	product2 := NewProduct("md", "Mouse", 6.7)
	product3 := NewProduct("lg", "Mac", 20.0)

	fmt.Println("Calculate Cost Product1")
	fmt.Println(product1.CalculateCost())
	fmt.Println("Calculate Cost Product2")
	fmt.Println(product2.CalculateCost())
	fmt.Println("Calculate Cost Product3")
	fmt.Println(product3.CalculateCost())

	newShop := NewShop(product1, product2)
	fmt.Println("Shop created")
	fmt.Println(newShop)
	newShop.Add(product3)
	fmt.Println("Product added in shop")
	fmt.Println(newShop)
	fmt.Println("Total")
	fmt.Println(newShop.Total())
}
