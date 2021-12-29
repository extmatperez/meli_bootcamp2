package main

import (
	"context"
	"fmt"

	"github.com/extmatperez/meli_bootcamp2/19_storage3/internal/product"
	"github.com/extmatperez/meli_bootcamp2/19_storage3/pkg/database"
)

func main() {
	repository := product.NewRepository(database.StorageDB)
	service := product.NewService(repository)

	products, err := service.GetAll(context.Background())

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(products)
}
