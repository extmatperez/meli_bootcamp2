package internal

import (
	"database/sql"
	"errors"
	"log"

	models "github.com/extmatperez/meli_bootcamp2/tree/castillo_walter/17_storage1/go-web-TM/internal/models"
	"github.com/extmatperez/meli_bootcamp2/tree/castillo_walter/17_storage1/go-web-TM/pkg/db"
)

type RepositorySQL interface {
	GetOne(id int) models.Product
	GetByName(nombre string) models.Product
	Store(product models.Product) (models.Product, error)
	Update(product models.Product) (models.Product, error)
}

type repositorySQL struct{}

func NewRepositorySQL() RepositorySQL {
	return &repositorySQL{}
}

func (r *repositorySQL) Store(product models.Product) (models.Product, error) {
	db := db.StorageDB

	stmt, err := db.Prepare("INSERT INTO producto(nombre, color, precio, stock, codigo, publicado, fecha_creacion) VALUES( ?, ?, ? ,? ,? , ?, ?)")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()
	var result sql.Result
	result, err = stmt.Exec(product.Nombre, product.Color, product.Precio, product.Stock, product.Codigo, product.Publicado, product.FechaCreacion)
	if err != nil {
		return models.Product{}, err
	}
	idCreado, _ := result.LastInsertId()
	product.ID = int(idCreado)

	return product, nil
}
func (r *repositorySQL) GetOne(id int) models.Product {
	db := db.StorageDB
	var productRead models.Product
	rows, err := db.Query("SELECT id, nombre,color FROM producto WHERE id = ?", id)

	if err != nil {
		log.Fatal(err)
		return productRead
	}

	for rows.Next() {
		err = rows.Scan(&productRead.ID, &productRead.Nombre, &productRead.Color) //, &productRead.Precio, &productRead.Stock, &productRead.Codigo, &productRead.Publicado, &productRead.FechaCreacion)
		if err != nil {
			log.Fatal(err)
			return productRead
		}
	}
	return productRead

}
func (r *repositorySQL) Update(product models.Product) (models.Product, error) {

	db := db.StorageDB

	stmt, err := db.Prepare("UPDATE producto SET nombre = ?, color = ?,precio = ? WHERE id = ?")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()
	result, err := stmt.Exec(product.Nombre, product.Color, product.Precio, product.ID)
	if err != nil {
		return models.Product{}, err
	}
	filasActualizadas, _ := result.RowsAffected()
	if filasActualizadas == 0 {
		return models.Product{}, errors.New("no se encontro el producto")
	}

	return product, nil

}

func (r *repositorySQL) GetByName(nombre string) models.Product {
	db := db.StorageDB
	var productRead models.Product
	rows, err := db.Query("SELECT id, nombre, color FROM producto WHERE nombre  = ?", nombre)

	if err != nil {
		log.Fatal(err)
		return productRead
	}

	for rows.Next() {
		err = rows.Scan(&productRead.ID, &productRead.Nombre, &productRead.Color) //, &productRead.Precio, &productRead.Stock, &productRead.Codigo, &productRead.Publicado, &productRead.FechaCreacion)
		if err != nil {
			log.Fatal(err)
			return productRead
		}
	}
	return productRead

}
