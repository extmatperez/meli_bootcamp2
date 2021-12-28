package internal

import (
	"log"

	"github.com/extmatperez/meli_bootcamp2/tree/scerca_nahuel/18_storage2/ProyectoEstructura/internal/models"
	"github.com/extmatperez/meli_bootcamp2/tree/scerca_nahuel/18_storage2/ProyectoEstructura/pkg/db"
)

type RepositorySQL interface {
	GetByName(name string) ([]models.Producto, error)
	Store(producto models.Producto) (models.Producto, error)
}

type repositorySQL struct {
}

func newRepositorySQL() RepositorySQL {
	return &repositorySQL{}
}

func (repsql *repositorySQL) Store(producto models.Producto) (models.Producto, error) {
	db := db.StorageDB

	stmt, err := db.Prepare("INSERT INTO productos(nombre,color,precio,stock,codigo,publicado,fechaCreacion) VALUES (?,?,?,?,?,?,?)")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	result, errExec := stmt.Exec(producto.Nombre, producto.Color, producto.Precio, producto.Stock, producto.Codigo, producto.Publicado, producto.FechaCreacion)
	if errExec != nil {
		return models.Producto{}, errExec
	}
	idCreado, _ := result.LastInsertId()
	producto.ID = int(idCreado)

	return producto, nil

}

func (r *repositorySQL) GetByName(name string) ([]models.Producto, error) {
	db := db.StorageDB
	var productosSlice []models.Producto
	var prodLeido models.Producto
	rows, err := db.Query("SELECT id, nombre,color,precio,stock,codigo,publicado,fechaCreacion FROM productos WHERE nombre = ?", name)

	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	for rows.Next() {
		err = rows.Scan(&prodLeido.ID, &prodLeido.Nombre, &prodLeido.Codigo, &prodLeido.Precio, &prodLeido.Stock, &prodLeido.Codigo, &prodLeido.Publicado, &prodLeido.FechaCreacion)
		if err != nil {
			log.Fatal(err)
			return nil, err
		}
		productosSlice = append(productosSlice, prodLeido)
	}
	return productosSlice, nil
}
