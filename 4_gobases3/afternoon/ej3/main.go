package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Product struct {
	name  string
	price float64
	count int
}

type Service struct {
	name    string
	price   float64
	minutes int
}

type Manteinance struct {
	name  string
	price float64
}

func sumProducts(products []Product, c chan float64) {
	var sum float64 = 0.0
	for _, p := range products {
		sum += p.price * float64(p.count)
	}
	fmt.Printf("La suma de productos es %.2f\n", sum)
	c <- sum
}

func sumServices(services []Service, c chan float64) {
	var sum float64 = 0.0
	var minutes int
	for _, s := range services {
		if s.minutes < 30 {
			minutes = 30
		} else {
			minutes = s.minutes
		}
		sum += s.price * float64(minutes)
	}
	fmt.Printf("La suma de servicios es %.2f\n", sum)
	c <- sum
}

func sumManteinance(manteinances []Manteinance, c chan float64) {
	var sum float64 = 0.0
	for _, m := range manteinances {
		sum += m.price
	}
	fmt.Printf("La suma de mantenimiento es %.2f\n", sum)
	c <- sum
}

func main() {
	rand.Seed(time.Now().UnixNano())

	var products []Product
	for i := 0; i < 10; i++ {
		price := (rand.Float64() * 100) + 1
		count := rand.Intn(3000)
		product := Product{"Alfajor", price, count}
		products = append(products, product)
	}

	var services []Service
	for i := 0; i < 10; i++ {
		price := (rand.Float64() * 10000) + 1
		minutes := rand.Intn(300)
		service := Service{"Internet", price, minutes}
		services = append(services, service)
	}

	var manteinances []Manteinance
	for i := 0; i < 10; i++ {
		price := (rand.Float64() * 1000) + 1
		manteinance := Manteinance{"Cleaning", price}
		manteinances = append(manteinances, manteinance)
	}

	c := make(chan float64)

	go sumProducts(products, c)
	go sumServices(services, c)
	go sumManteinance(manteinances, c)

	var sum float64 = 0.0
	for i := 0; i < 3; i++ {
		sum += <-c
	}

	fmt.Printf("La suma de todo es: %.2f\n", sum)
}
