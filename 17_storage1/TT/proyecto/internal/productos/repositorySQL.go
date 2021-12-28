/*
Desarrollar un método en el repositorio que permita hacer búsquedas de un producto por nombre. Para lograrlo se deberá:
	- Diseñar interfaz “Repository” en la que exista un método GetByName() que reciba por parámetro un string y
	retorne un objeto del tipo Product.
	- Implementar el método de forma que con el string recibido lo use para buscar en la DB por el campo “name”.
*/

package internal

import (
	"database/sql"
	"errors"
	"log"

	"github.com/extmatperez/meli_bootcamp2/17_storage1/TT/proyecto/internal/models"
	"github.com/extmatperez/meli_bootcamp2/17_storage1/TT/proyecto/pkg/db"

	_ "github.com/go-sql-driver/mysql"
)

type RepositorySQL interface {
	Store(nuevoProducto models.Producto) (models.Producto, error)
	Get(id int) (models.Producto, error)
	GetByName(nombre string) ([]models.Producto, error)
	Update(productoToUpdate models.Producto) (models.Producto, error)
}

type repositorySQL struct{}

func NewRepositorySQL() RepositorySQL {
	return &repositorySQL{}
}

func (r *repositorySQL) Store(nuevoProducto models.Producto) (models.Producto, error) {

	db := db.StorageDB
	statement, err := db.Prepare("insert into productos(nombre, color, stock, precio, codigo, publicado, fecha_creacion) values(?,?,?,?,?,?,?);")

	if err != nil {
		log.Fatal(err)
	}

	defer statement.Close()

	result, err := statement.Exec(nuevoProducto.Nombre, nuevoProducto.Color, nuevoProducto.Stock, nuevoProducto.Precio, nuevoProducto.Codigo, nuevoProducto.Publicado, nuevoProducto.FechaCreacion)

	if err != nil {
		return models.Producto{}, err
	}

	idCreado, _ := result.LastInsertId()

	nuevoProducto.ID = int(idCreado)

	return nuevoProducto, nil
}

func (r *repositorySQL) Get(id int) (models.Producto, error) {

	db := db.StorageDB
	rows, err := db.Query("select id, nombre, color, precio, stock, codigo, publicado, fecha_creacion from productos where id = ?;", id)

	if err != nil {
		log.Fatal(err)
		return models.Producto{}, err
	}

	var productoLeido models.Producto

	for rows.Next() {
		err := rows.Scan(&productoLeido.ID, &productoLeido.Nombre, &productoLeido.Color, &productoLeido.Precio, &productoLeido.Stock, &productoLeido.Codigo, &productoLeido.Publicado, &productoLeido.FechaCreacion)

		if err != nil {
			log.Fatal(err)
			return productoLeido, err
		}
	}

	return productoLeido, nil
}

func (r *repositorySQL) GetByName(nombre string) ([]models.Producto, error) {

	db := db.StorageDB
	rows, err := db.Query("select id, nombre, color, precio, stock, codigo, publicado, fecha_creacion from productos where nombre like ?;", nombre)

	if err != nil {
		log.Fatal(err)
		return []models.Producto{}, err
	}

	var productoLeido models.Producto
	var productos []models.Producto

	for rows.Next() {
		err := rows.Scan(&productoLeido.ID, &productoLeido.Nombre, &productoLeido.Color, &productoLeido.Precio, &productoLeido.Stock, &productoLeido.Codigo, &productoLeido.Publicado, &productoLeido.FechaCreacion)

		if err != nil {
			log.Fatal(err)
			return []models.Producto{}, err
		}

		productos = append(productos, productoLeido)
	}

	return productos, nil
}

func (r *repositorySQL) Update(productoToUpdate models.Producto) (models.Producto, error) {

	db := db.StorageDB
	statement, err := db.Prepare("update productos set nombre = ?, color = ?, stock = ?, precio = ?, codigo = ?, publicado = ?, fecha_creacion = ? where id = ?;")

	if err != nil {
		log.Fatal(err)
	}

	defer statement.Close()

	var result sql.Result

	result, err = statement.Exec(productoToUpdate.Nombre, productoToUpdate.Color, productoToUpdate.Stock, productoToUpdate.Precio, productoToUpdate.Codigo, productoToUpdate.Publicado, productoToUpdate.FechaCreacion, productoToUpdate.ID)

	if err != nil {
		return models.Producto{}, err
	}

	rowsAffected, err := result.RowsAffected()

	if err != nil {
		log.Fatal(err)
	}

	if rowsAffected == 0 {
		return models.Producto{}, errors.New("no se encontró el producto")
	}

	return productoToUpdate, nil
}
