package internal

import (
	"database/sql"
	"log"
	"strings"

	"github.com/extmatperez/meli_bootcamp2/17_storage1/TT/go-web/internal/models"
	"github.com/extmatperez/meli_bootcamp2/17_storage1/TT/go-web/pkg/db"
)

type RepositorySQL interface {
	GetByName(name string) (models.Product, error)
	Store(nombre string, color string, precio int, stock int, codigo string, publicado bool, fechaCreacion string) (models.Product, error)
}

type repositorySQL struct{}

func NewRepositorySQL() RepositorySQL {
	return &repositorySQL{}
}

func (r *repositorySQL) GetByName(name string) (models.Product, error) {
	db := db.StorageDB
	var productReaded models.Product
	rows, err := db.Query("select id, nombre, color, precio, stock, codigo, publicado, fechaCreacion from products where LOWER(nombre) = ?", strings.ToLower(name))

	if err != nil {
		log.Fatal(err)
		return productReaded, err
	}

	for rows.Next() {
		err = rows.Scan(&productReaded.ID, &productReaded.Nombre, &productReaded.Color, &productReaded.Precio, &productReaded.Stock, &productReaded.Codigo, &productReaded.Publicado, &productReaded.FechaCreacion)
		if err != nil {
			log.Fatal(err)
			return productReaded, err
		}
	}
	return productReaded, nil
}

func (r *repositorySQL) Store(nombre string, color string, precio int, stock int, codigo string, publicado bool, fechaCreacion string) (models.Product, error) {
	db := db.StorageDB

	stmt, err := db.Prepare("INSERT INTO products(nombre, color, precio, stock, codigo, publicado, fechaCreacion) VALUES( ?, ?, ? , ? , ? ,?, ?)")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()
	var result sql.Result
	result, err = stmt.Exec(nombre, color, precio, stock, codigo, publicado, fechaCreacion)
	if err != nil {
		return models.Product{}, err
	}
	idCreado, _ := result.LastInsertId()
	productCreated := models.Product{ID: int(idCreado), Nombre: nombre, Color: color, Precio: precio, Stock: stock, Codigo: codigo, Publicado: publicado, FechaCreacion: fechaCreacion}

	return productCreated, nil

}
