package main

import (
	"context"
	"fmt"

	"github.com/extmatperez/meli_bootcamp2/18_storage2/internal/product"
	"github.com/extmatperez/meli_bootcamp2/18_storage2/pkg/database"
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
