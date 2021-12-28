package internal

import (
	"database/sql"
	"log"

	"github.com/extmatperez/meli_bootcamp2/17_storage1/proyecto/internal/models"
	"github.com/extmatperez/meli_bootcamp2/17_storage1/proyecto/pkg/db"
)

type RepositorySql interface {
	Store(prod models.Productos) (models.Productos, error)
	GetByName(name string) models.Productos
}

type repositorySql struct{}

func NewRepositorySql() RepositorySql {
	return &repositorySql{}
}

func (s *repositorySql) Store(prod models.Productos) (models.Productos, error) {
	db := db.StorageDB
	stmt, err := db.Prepare("INSERT INTO productos(nombre, color, precio, stock, codigo, publicado, fecha_de_creacion) VALUES(?,?,?,?,?,?,?) ")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()
	var result sql.Result
	result, err = stmt.Exec(prod.Nombre, prod.Color, prod.Precio, prod.Stock, prod.Codigo, prod.Publicado, prod.Fecha_de_creacion)
	if err != nil {
		return models.Productos{}, err
	}
	idCreado, _ := result.LastInsertId()
	prod.Id = int(idCreado)
	return prod, nil
}

func (s *repositorySql) GetByName(name string) models.Productos {
	db := db.StorageDB
	var prodBuscado models.Productos
	rows, err := db.Query("SELECT id, nombre, color, stock, publicado, fecha_de_creacion FROM productos WHERE nombre = ? ", name)
	if err != nil {
		log.Fatal(err)
		return prodBuscado
	}
	for rows.Next() {
		err = rows.Scan(&prodBuscado.Id, &prodBuscado.Nombre, &prodBuscado.Color, &prodBuscado.Stock, &prodBuscado.Publicado, &prodBuscado.Fecha_de_creacion)
		if err != nil {
			log.Fatal(err)
			return prodBuscado
		}
	}
	return prodBuscado
}
