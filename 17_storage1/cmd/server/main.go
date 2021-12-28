package main

import (
	"context"
	"fmt"

	"github.com/extmatperez/meli_bootcamp2/17_storage1/internal/product"
)

func main() {
	repository := product.NewRepository()
	service := product.NewService(repository)

	products, err := service.GetAll(context.Background())

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(products)
}
