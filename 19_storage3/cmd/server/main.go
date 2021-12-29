package main

import (
	"context"
	"fmt"
	"log"

	"github.com/extmatperez/meli_bootcamp2/19_storage3/internal/product"
	"github.com/extmatperez/meli_bootcamp2/19_storage3/pkg/database"
)

func main() {
	// DYNAMO DB
	dynamoDB, err := database.InitDynamo()
	if err != nil {
		log.Fatal(err)
	}
	repository := product.NewDynamoRepository(dynamoDB, "products")

	products, _ := repository.GetAll(context.Background())

	fmt.Println("prod: begin")
	fmt.Println(products)
	fmt.Println("prod: end")

	// MYSQL DB
	// repository := product.NewRepository(database.StorageDB)
	// service := product.NewService(repository)

	// products, err := service.GetAll(context.Background())

	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }

	// fmt.Println(products)
}
