package main

import (
	"context"
	"fmt"

	"github.com/extmatperez/meli_bootcamp2/19_storage3/internal/product"
	"github.com/extmatperez/meli_bootcamp2/19_storage3/pkg/database"
)

func main() {
	// MONGO DB
	repository := product.NewMongoRepository(database.MongoDB)
	_ = repository.SetDatabaseAndCollection("bootcamp_storage", "products")
	products, err := repository.GetAll(context.Background())

	if err != nil {
		fmt.Println("errr")
		fmt.Println(err)
	}

	fmt.Println("prod: begin")
	fmt.Println(products)
	fmt.Println("prod: end")

	// product, err := repository.Store(context.Background(), domain.ProductMongo{
	// 	Name:        "Prueba",
	// 	Price:       1.1,
	// 	Description: "Description",
	// })

	// fmt.Println(product)

	// DYNAMO DB
	// dynamoDB, err := database.InitDynamo()
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// repository := product.NewDynamoRepository(dynamoDB, "products")

	// products, _ := repository.GetAll(context.Background())

	// fmt.Println("prod: begin")
	// fmt.Println(products)
	// fmt.Println("prod: end")

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
