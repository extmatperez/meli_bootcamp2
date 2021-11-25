package main

import (
	"fmt"
)

func main() {
	product1 := Product{
		Name:   "Product1",
		Price:  150.00,
		Amount: 10,
	}
	product2 := Product{
		Name:   "Product2",
		Price:  1500.50,
		Amount: 5,
	}
	product3 := Product{
		Name:   "Product3",
		Price:  13000.00,
		Amount: 2,
	}

	service1 := Service{
		Name:          "Service1",
		Price:         3000.00,
		MinutesWorked: 300,
	}
	service2 := Service{
		Name:          "Service2",
		Price:         1500.00,
		MinutesWorked: 25,
	}

	maintenance1 := Maintenance{
		Name:  "Maintenance1",
		Price: 50.0,
	}
	maintenance2 := Maintenance{
		Name:  "Maintenance2",
		Price: 25.0,
	}

	var finalPrice float64

	productsPriceChannel := make(chan float64)
	servicesPriceChannel := make(chan float64)
	maintainancesPriceChannel := make(chan float64)

	go CalculatePriceOfProducts(productsPriceChannel, product1, product2, product3)
	go CalculatePriceOfServices(servicesPriceChannel, service1, service2)
	go CalculatePriceOfMaintainances(maintainancesPriceChannel, maintenance1, maintenance2)

	finalPrice = <-productsPriceChannel + <-servicesPriceChannel + <-maintainancesPriceChannel
	fmt.Printf("Precio total: %.2f\n", finalPrice)
}

type Product struct {
	Name   string
	Price  float64
	Amount int
}

type Service struct {
	Name          string
	Price         float64
	MinutesWorked int
}

type Maintenance struct {
	Name  string
	Price float64
}

func CalculatePriceOfProducts(priceOfProducts chan float64, products ...Product) {
	var total float64

	for _, product := range products {
		total += product.Price * float64(product.Amount)
	}

	priceOfProducts <- total
}

func CalculatePriceOfServices(priceOfServices chan float64, services ...Service) {
	var total float64

	for _, service := range services {
		minutesToCharge := int(service.MinutesWorked / 30)

		if minutesToCharge == 0 {
			minutesToCharge = 1
		}

		total += service.Price * float64(minutesToCharge)
	}
	priceOfServices <- total
}

func CalculatePriceOfMaintainances(priceOfMaintainances chan float64, maintenances ...Maintenance) {
	var total float64

	for _, maintenance := range maintenances {
		total += maintenance.Price
	}
	priceOfMaintainances <- total
}
