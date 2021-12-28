package internal

import (
	"database/sql"
	"log"

	"github.com/extmatperez/meli_bootcamp2/17_storage1/afternoon/go-web/internal/models"
	"github.com/extmatperez/meli_bootcamp2/17_storage1/afternoon/go-web/pkg/db"
)

type RepositorySQL interface {
	Store(product models.Producto) (models.Producto, error)
	GetByName(name string) ([]models.Producto, error)
}

type repositorySQL struct{}

func NewRepositorySQL() RepositorySQL {
	return &repositorySQL{}
}

func (repo *repositorySQL) Store(producto models.Producto) (models.Producto, error) {
	db := db.StorageDB                                                                                                                             // se inicializa la base
	stmt, err := db.Prepare("INSERT INTO productos(nombre, color, precio, stock, codigo, publicado, fechaCreacion) VALUES( ?, ?, ?, ?, ?, ?, ? )") // se prepara el SQL
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close() // se cierra la sentencia al terminar. Si quedan abiertas se genera consumos de memoria
	var result sql.Result
	result, err = stmt.Exec(producto.Nombre, producto.Color, producto.Precio, producto.Stock, producto.Codigo, producto.Publicado, producto.FechaCreacion) // retorna un sql.Result y un error
	if err != nil {
		return models.Producto{}, err
	}
	insertedId, _ := result.LastInsertId() // del sql.Resul devuelto en la ejecución obtenemos el Id insertado
	producto.ID = int(insertedId)
	return producto, nil
}

func (repo *repositorySQL) GetByName(name string) ([]models.Producto, error) {
	var producto models.Producto
	var productos []models.Producto
	db := db.StorageDB

	rows, err := db.Query("select * from productos where nombre = ?", name)
	if err != nil {
		log.Println(err)
		return []models.Producto{}, err
	}
	for rows.Next() {
		if err := rows.Scan(&producto.ID, &producto.Nombre, &producto.Color, &producto.Precio, &producto.Stock, &producto.Codigo, &producto.Publicado, &producto.FechaCreacion); err != nil {
			log.Println(err.Error())
			return []models.Producto{}, err
		}
		productos = append(productos, producto)
	}
	return productos, nil
}
