package main

import (
	internal "github.com/extmatperez/meli_bootcamp2/tree/pescie_juan/hackathon/internal/productos"
)

func main() {

	repository := internal.NewRepositorySQL()
	service := internal.NewServiceSQL(repository)
	service.LoadProductsOnDB()

}
