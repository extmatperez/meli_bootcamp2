package main

import "fmt"

const (
	SMALL = "small"
	MEDIUM = "medium"
	BIG = "big"
)

type Product struct {
	p_type string
	name string
	price float64
}

type Store struct {
	products []IProduct
}

func (store *Store) AddProduct(p IProduct) {
	store.products = append(store.products, p)
}

func (store Store) TotalCost() float64 {
	var total float64
	for _, p := range store.products {
		total += p.CalcularCosto()
	}
	return total
}

type IProduct interface {
	CalcularCosto() float64
}

type IEcommerce interface {
	AddProduct(IProduct)
	TotalCost() float64
}

func newProduct(p_type string, name string, price float64) IProduct {
	if (p_type == SMALL) {
		return SmallProduct{Product{p_type, name, price}}
	} else if (p_type == BIG) {
		return BigProduct{Product{p_type, name, price}}
	} else if (p_type == MEDIUM) {
		return MediumProduct{Product{p_type, name, price}}
	} else {
		return nil
	}
}

func newStore() IEcommerce {
	return &Store{}
}

type SmallProduct struct {
	p Product
}

func (p SmallProduct) CalcularCosto() float64 {
	return p.p.price
}

type MediumProduct struct {
	p Product
}

func (p MediumProduct) CalcularCosto() float64 {
	return p.p.price * 1.03
}

type BigProduct struct {
	p Product
}

func (p BigProduct) CalcularCosto() float64 {
	return p.p.price * 1.06 + 2500
}

func main() {
	store := newStore()
	store.AddProduct(newProduct(SMALL, "Laptop", 1500))
	store.AddProduct(newProduct(MEDIUM, "Tablet", 500))
	store.AddProduct(newProduct(BIG, "Desktop", 5000))
	fmt.Println(store.TotalCost())
}