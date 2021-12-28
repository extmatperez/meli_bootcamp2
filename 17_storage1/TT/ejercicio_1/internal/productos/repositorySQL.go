package internal

import (
	"database/sql"
	"log"

	"github.com/extmatperez/meli_bootcamp2/tree/parra_diego/17_storage1/TT/ejercicio_1/internal/models"
	"github.com/extmatperez/meli_bootcamp2/tree/parra_diego/17_storage1/TT/ejercicio_1/pkg/db"
	_ "github.com/go-sql-driver/mysql"
)

type RepositorySQL interface {
	Store(persona models.Product) (models.Product, error)
	GetOneName(name string) models.Product
	// Update(persona models.Product) (models.Product, error)
}
type repositorySQL struct{}

func NewRepositorySQL() RepositorySQL {
	return &repositorySQL{}
}
func (r *repositorySQL) Store(product models.Product) (models.Product, error) {
	db := db.StorageDB

	stmt, err := db.Prepare("INSERT INTO product (name, color, price) VALUES( ?, ?, ? )")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()
	var result sql.Result
	result, err = stmt.Exec(product.Name, product.Color, product.Price)
	if err != nil {
		return models.Product{}, err
	}
	idCreado, _ := result.LastInsertId()
	product.Id = int(idCreado)

	return product, nil
}
func (r *repositorySQL) GetOneName(name string) models.Product {
	db := db.StorageDB
	var productoLeido models.Product
	rows, err := db.Query("SELECT id, name,color, price FROM product WHERE name like ?", name)

	if err != nil {
		log.Fatal(err)
		return productoLeido
	}

	for rows.Next() {
		err = rows.Scan(&productoLeido.Id, &productoLeido.Name, &productoLeido.Color, &productoLeido.Price)
		if err != nil {
			log.Fatal(err)
			return productoLeido
		}
	}
	return productoLeido
}
