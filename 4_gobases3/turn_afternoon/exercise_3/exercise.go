package main

import (
	"fmt"
	"math"
)

func main() {

	products := []*Product{}
	product := &Product{"Apple", 100.50, 5}
	product2 := &Product{"Strawberry", 500.50, 5}
	products = append(products, product, product2)
	fmt.Println("----Total Products:----")
	fmt.Println(pricesProductResulter(products))

	services := []*Service{}
	service := &Service{"Door", 900.50, 180}
	service2 := &Service{"Floor", 800.50, 220}
	services = append(services, service, service2)
	fmt.Println("----Total Services:----")
	fmt.Println(pricesServiceResulter(services))

	maintenances := []*Maintenance{}
	maintenance := &Maintenance{"Car wheel", 1500.00}
	maintenance2 := &Maintenance{"Car engine", 2000.00}
	maintenances = append(maintenances, maintenance, maintenance2)
	fmt.Println("----Total Maintenances:----")
	fmt.Println(pricesMaintenanceResulter(maintenances))

}

type Product struct {
	name     string
	price    float64
	quantity int
}
type Service struct {
	name          string
	price         float64
	minutesWorked int
}
type Maintenance struct {
	name  string
	price float64
}

func pricesProductResulter(prods []*Product) float64 {
	var total float64
	for _, prod := range prods {
		total += prod.price
	}
	return total
}

func pricesServiceResulter(services []*Service) float64 {
	var totalPrice float64
	var total30Minutes float64
	for _, service := range services {
		total30Minutes += math.RoundToEven(float64(service.minutesWorked) / 30)
		totalPrice += service.price
	}
	totalPriceByMinutes := totalPrice * total30Minutes
	return totalPriceByMinutes
}

func pricesMaintenanceResulter(maints []*Maintenance) float64 {
	var total float64
	for _, main := range maints {
		total += main.price
	}
	return total
}
