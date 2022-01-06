package internal

import (
	"errors"
	"log"
	"strings"

	"github.com/extmatperez/meli_bootcamp2/tree/vega_rodrigo/hackaton/internal/models"
	"github.com/extmatperez/meli_bootcamp2/tree/vega_rodrigo/hackaton/pkg/db"
	"github.com/extmatperez/meli_bootcamp2/tree/vega_rodrigo/hackaton/pkg/store"
)

type ProductRepository interface {
	ImportAllProducts() error
	StoreProduct(models.Product) (models.Product, error)
	UpdateProduct(models.Product) (models.Product, error)
}

type repository_product struct {
	arr store.SaveFile
}

func NewProductRepository(arr store.SaveFile) ProductRepository {
	return &repository_product{arr}
}

func (r *repository_product) ImportAllProducts() error {
	products_string, err := r.arr.ReadLines("/Users/rovega/Documents/GitHub/meli_bootcamp2/hackaton/cmd/server/data/products.txt")
	if err != nil {
		return err
	}

	for _, product := range products_string {
		only_product := strings.Split(product, "#$%#")
		id := only_product[0]
		description := only_product[1]
		price := only_product[2]

		db := db.StorageDB
		query := "INSERT INTO Product(id, `description`, price) VALUES (?,?,?)"
		stmt, err := db.Prepare(query)
		if err != nil {
			log.Fatal(err)
		}
		defer stmt.Close()

		_, err = stmt.Exec(id, description, price)

		if err != nil {
			return errors.New("No se pudo guardar elemento en BD.")
		}
	}
	return nil
}

func (r *repository_product) StoreProduct(product models.Product) (models.Product, error) {
	db := db.StorageDB
	query := "INSERT INTO Product(`description`, price) VALUES (?,?)"
	stmt, err := db.Prepare(query)
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	result, err := stmt.Exec(product.Description, product.Price)
	if err != nil {
		return models.Product{}, err
	}

	idCreado, _ := result.LastInsertId()
	product.Id = int(idCreado)
	return product, nil
}

func (r *repository_product) UpdateProduct(product models.Product) (models.Product, error) {
	db := db.StorageDB
	query := "UPDATE Product SET description = ?, price = ? WHERE id = ?"
	stmt, err := db.Prepare(query)
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()
	result, err := stmt.Exec(product.Description, product.Price, product.Id)
	if err != nil {
		return models.Product{}, err
	}
	updatedRows, _ := result.RowsAffected()
	if updatedRows == 0 {
		return models.Product{}, errors.New("No se encontro al product.")
	}
	return product, nil
}
