package main

import "fmt"

const (
	small  = "sm"
	medium = "md"
	large  = "bg"
)

type Ecommerce interface {
	Total() float64
	Add(p Product)
}

type store struct {
	Products []Product
}

func NewStore() Ecommerce {
	var products []Product = make([]Product, 0)
	var sto = store{products}
	var ecom Ecommerce = &sto
	return ecom
}

func (s *store) Add(p Product) {
	s.Products = append(s.Products, p)
}

func (s store) Total() float64 {
	return 0.0
}

type Product interface {
	CostCalculation() float64
}

type product struct {
	TypeN string
	Name  string
	Price float64
}

func (p product) CostCalculation() float64 {
	var typesCalc = map[string]float64{
		"sm": p.Price,
		"md": p.Price * 0.03,
		"bg": (p.Price * 0.06) + 2500,
	}

	return typesCalc[p.TypeN]
}

func NewProduct(typeN string, name string, price float64) Product {
	return product{TypeN: typeN, Name: name, Price: price}
}

func main() {
	var sto = NewStore()

	types := []string{small, medium, large}

	for i := 0; i < 10; i++ {
		p := NewProduct(types[i], "Alfajor", float64(i*i))
		sto.Add(p)
	}

	fmt.Println(sto.Total())
}
