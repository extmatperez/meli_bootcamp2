/*
Una empresa nacional se encarga de realizar venta de productos, servicios y mantenimiento.
Para ello requieren realizar un programa que se encargue de calcular el precio total de Productos, Servicios y Mantenimientos.
Debido a la fuerte demanda y para optimizar la velocidad requieren que el cálculo de la sumatoria se realice en paralelo
mediante 3 go routines.

Se requieren 3 estructuras:
Productos: nombre, precio, cantidad.
Servicios: nombre, precio, minutos trabajados.
Mantenimiento: nombre, precio.

Se requieren 3 funciones:
Sumar Productos: recibe un array de producto y devuelve el precio total (precio * cantidad).
Sumar Servicios: recibe un array de servicio y devuelve el precio total (precio * media hora trabajada, si no llega a trabajar
	30 minutos se le cobra como si hubiese trabajado media hora).
Sumar Mantenimiento: recibe un array de mantenimiento y devuelve el precio total.

Los 3 se deben ejecutar concurrentemente y al final se debe mostrar por pantalla el monto final (sumando el total de los 3).


*/

package main

import (
	"fmt"
)

type Product struct {
	Name     string
	Price    float64
	Quantity int
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

func addProduct(priceProductsTotals chan float64, arrayProduct []Product) {
	fmt.Println("Inicia add Product")
	var totalPrice float64 = 0.0
	for _, prod := range arrayProduct {
		totalPrice += (prod.Price * float64(prod.Quantity))
	}
	fmt.Println("Finaliza add Product, total: ", totalPrice)
	priceProductsTotals <- totalPrice
}

func addServices(priceServicesTotals chan float64, arrayService []Service) {
	fmt.Println("Inicia add Service")
	var totalPrice float64 = 0.0
	for _, serv := range arrayService {
		if serv.MinutesWorked < 30 {
			totalPrice += (30.0 * (serv.Price))
		} else {
			totalPrice += (serv.Price * float64(serv.MinutesWorked))
		}
	}
	fmt.Println("Finaliza add Service, total: ", totalPrice)
	priceServicesTotals <- totalPrice
}

func addMaintenance(priceMaintenanceTotals chan float64, arrayMaintenance []Maintenance) {
	fmt.Println("Inicia add Maintenance")
	var totalPrice float64 = 0.0
	for _, maint := range arrayMaintenance {
		totalPrice += maint.Price
	}
	fmt.Println("Finaliza add Maintenance, total: ", totalPrice)
	priceMaintenanceTotals <- totalPrice
}
func main() {
	var listProduct []Product
	var listServices []Service
	var listMaintenance []Maintenance

	prod1 := Product{"Mouse", 12.85, 5}
	prod2 := Product{"Teclado", 105.42, 2}
	prod3 := Product{"Monitor", 110.22, 1}
	listProduct = append(listProduct, prod1)
	listProduct = append(listProduct, prod2)
	listProduct = append(listProduct, prod3)
	priceProductsTotals := make(chan float64)

	serv1 := Service{"ArmadoEscritorio", 12.85, 5}
	serv2 := Service{"Instalar SO", 5.42, 2}
	listServices = append(listServices, serv1)
	listServices = append(listServices, serv2)
	priceServicesTotals := make(chan float64)

	maint1 := Maintenance{"Retrosoplado", 12.85}
	maint2 := Maintenance{"Limpieza de ordenador", 5.42}
	maint3 := Maintenance{"Ayuda", 5.22}
	listMaintenance = append(listMaintenance, maint1)
	listMaintenance = append(listMaintenance, maint2)
	listMaintenance = append(listMaintenance, maint3)
	priceMaintenanceTotals := make(chan float64)

	var totalFinal float64 = 0.0

	go addProduct(priceProductsTotals, listProduct)
	go addServices(priceServicesTotals, listServices)
	go addMaintenance(priceMaintenanceTotals, listMaintenance)

	totalFinal = <-priceProductsTotals + <-priceServicesTotals + <-priceMaintenanceTotals
	fmt.Println(totalFinal)

}
