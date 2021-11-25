package main

import "fmt"

type Product struct {
	Name   string
	Price  float64
	Amount int
}

type Service struct {
	Name    string
	Price   float64
	Minutes int
}

type Mainteinance struct {
	Name  string
	Price float64
}

func accProducts(products []Product, c chan float64) {
	acc := 0.0
	for _, product := range products {
		acc += product.Price * float64(product.Amount)
	}
	fmt.Println("p: ", acc)
	c <- acc
}

func accServices(services []Service, c chan float64) {
	acc := 0.0
	for _, service := range services {
		acc += service.Price * float64(service.Minutes/30)
	}
	fmt.Println("s: ", acc)
	c <- acc
}

func accMainteinance(maints []Mainteinance, c chan float64) {
	acc := 0.0
	for _, maint := range maints {
		acc += maint.Price
	}
	fmt.Println("m: ", acc)
	c <- acc
}

func main() {
	products := []Product{{Name: "tomate", Price: 3.00, Amount: 5}, {Name: "zanahoria", Price: 2.50, Amount: 3}}
	services := []Service{{Name: "limpieza", Price: 5.00, Minutes: 130}, {Name: "formateo", Price: 8.00, Minutes: 320}}
	maints := []Mainteinance{{Name: "mantenimiento1", Price: 15.50}, {Name: "mantenimiento2", Price: 12.00}}

	c := make(chan float64)
	acc := 0.0

	go accProducts(products, c)
	acc += <-c
	go accServices(services, c)
	acc += <-c
	go accMainteinance(maints, c)
	acc += <-c

	fmt.Println("Total: ", acc)
}
