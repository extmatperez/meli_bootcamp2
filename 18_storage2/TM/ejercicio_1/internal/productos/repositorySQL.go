package internal

import (
	"context"
	"database/sql"
	"errors"
	"log"

	"github.com/extmatperez/meli_bootcamp2/tree/parra_diego/18_storage2/TM/ejercicio_1/internal/models"
	"github.com/extmatperez/meli_bootcamp2/tree/parra_diego/18_storage2/TM/ejercicio_1/pkg/db"
	_ "github.com/go-sql-driver/mysql"
)

type RepositorySQL interface {
	Store(product models.Product) (models.Product, error)
	GetOneName(name string) ([]models.Product, error)
	GetAll() ([]models.Product, error)
	Update(ctx context.Context, product models.Product) (models.Product, error)
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
func (r *repositorySQL) GetOneName(name string) ([]models.Product, error) {
	db := db.StorageDB
	var productoLeido models.Product
	var listProducts []models.Product
	rows, err := db.Query("SELECT id, name,color, price FROM product WHERE name = ?", name)

	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	for rows.Next() {
		err = rows.Scan(&productoLeido.Id, &productoLeido.Name, &productoLeido.Color, &productoLeido.Price)
		if err != nil {
			log.Fatal(err)
			return nil, err
		}
		listProducts = append(listProducts, productoLeido)
	}
	return listProducts, nil
}
func (r *repositorySQL) GetAll() ([]models.Product, error) {
	db := db.StorageDB
	var productoLeido models.Product
	var listProducts []models.Product
	rows, err := db.Query("SELECT id, name,color, price FROM product")

	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	for rows.Next() {
		err = rows.Scan(&productoLeido.Id, &productoLeido.Name, &productoLeido.Color, &productoLeido.Price)
		if err != nil {
			log.Fatal(err)
			return nil, err
		}
		listProducts = append(listProducts, productoLeido)
	}
	return listProducts, nil
}
func (r *repositorySQL) Update(ctx context.Context, product models.Product) (models.Product, error) {
	db := db.StorageDB

	stmt, err := db.PrepareContext(ctx, "UPDATE product SET name=?,color=?,price=? WHERE id = ?")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()
	var result sql.Result
	result, err = stmt.ExecContext(ctx, product.Name, product.Color, product.Price, product.Id)
	if err != nil {
		return models.Product{}, err
	}
	filasActualizadas, _ := result.RowsAffected()
	if filasActualizadas == 0 {
		return models.Product{}, errors.New("No se encontro el producto")
	}

	return product, nil
}
