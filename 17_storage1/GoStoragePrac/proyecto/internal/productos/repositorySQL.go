package internal

import (
	"database/sql"
	"errors"
	"log"

	"github.com/extmatperez/meli_bootcamp2/17_storage1/GoStoragePrac/proyecto/internal/models"
	"github.com/extmatperez/meli_bootcamp2/17_storage1/GoStoragePrac/proyecto/pkg/db"
)

type RepositorySQL interface {
	Store(producto models.Producto) (models.Producto, error)
	GetOne(id int) models.Producto
	GetAll() ([]models.Producto, error)
	Update(producto models.Producto) (models.Producto, error)
	Delete(id int) error
	GetByName(nombre string) []models.Producto
}

type repositorySQL struct {
}

func NewRepositorySQL() RepositorySQL {
	return &repositorySQL{}
}

func (repo *repositorySQL) Store(producto models.Producto) (models.Producto, error) {

	db := db.StorageDB
	stmt, err := db.Prepare("INSERT INTO products(nombre, color, precio, stock, codigo, publicado, fechaCreacion) VALUES( ?, ?, ?, ?, ?, ?, ?)") // se prepara el SQL
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()
	var result sql.Result
	result, err = stmt.Exec(producto.Nombre, producto.Color, producto.Precio, producto.Stock, producto.Codigo, producto.Publicado, producto.FechaCreacion)
	if err != nil {
		return models.Producto{}, err
	}
	insertedId, _ := result.LastInsertId()
	producto.Id = int(insertedId)
	return producto, nil

}

func (repo *repositorySQL) GetOne(id int) models.Producto {

	db := db.StorageDB
	var productoLeido models.Producto
	rows, err := db.Query("SELECT id, nombre, color, precio, stock, codigo, publicado, fechaCreacion FROM products WHERE id = ?", id)

	if err != nil {
		log.Fatal(err)
		return productoLeido
	}

	for rows.Next() {
		err = rows.Scan(&productoLeido.Id, &productoLeido.Nombre, &productoLeido.Color, &productoLeido.Precio, &productoLeido.Stock, &productoLeido.Codigo, &productoLeido.Publicado, &productoLeido.FechaCreacion)
		if err != nil {
			log.Fatal(err)
			return productoLeido
		}
	}
	return productoLeido

}

func (repo *repositorySQL) GetAll() ([]models.Producto, error) {
	var misProductos []models.Producto
	db := db.StorageDB
	var productoLeido models.Producto
	rows, err := db.Query("SELECT * FROM products")
	if err != nil {
		log.Println(err)
		return nil, err
	}

	for rows.Next() {

		if err := rows.Scan(&productoLeido.Id, &productoLeido.Nombre, &productoLeido.Color, &productoLeido.Precio, &productoLeido.Stock, &productoLeido.Codigo, &productoLeido.Publicado, &productoLeido.FechaCreacion); err != nil {
			log.Fatal(err)
			return nil, err
		}

		//APPENDEAMOS A PRODUCTS
		misProductos = append(misProductos, productoLeido)
	}
	return misProductos, nil

}
func (repo *repositorySQL) Update(producto models.Producto) (models.Producto, error) {

	db := db.StorageDB

	stmt, err := db.Prepare("UPDATE products SET nombre = ?, color = ?, stock = ? WHERE id = ?")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()
	result, err := stmt.Exec(producto.Nombre, producto.Color, producto.Stock, producto.Id)
	if err != nil {
		return models.Producto{}, err
	}
	filasActualizadas, _ := result.RowsAffected()
	if filasActualizadas == 0 {
		return models.Producto{}, errors.New("no se encontro el producto")
	}

	return producto, nil

}

func (repo *repositorySQL) Delete(id int) error {

	db := db.StorageDB
	stmt, err := db.Prepare("DELETE FROM products WHERE id = ?")
	if err != nil {
		log.Fatal(err)
	}

	defer stmt.Close()
	result, err := stmt.Exec(id)
	if err != nil {
		return err
	}
	filasActualizadas, _ := result.RowsAffected()
	if filasActualizadas == 0 {
		return errors.New("no se encontro el producto")
	}

	return nil

}

func (repo *repositorySQL) GetByName(nombre string) []models.Producto {

	db := db.StorageDB
	var productoLeido models.Producto
	rows, err := db.Query("SELECT id, nombre, color, precio, stock, codigo, publicado, fechaCreacion FROM products WHERE nombre = ?", nombre)
	var sliceProductos []models.Producto
	if err != nil {
		log.Fatal(err)
		return sliceProductos
	}

	for rows.Next() {
		err = rows.Scan(&productoLeido.Id, &productoLeido.Nombre, &productoLeido.Color, &productoLeido.Precio, &productoLeido.Stock, &productoLeido.Codigo, &productoLeido.Publicado, &productoLeido.FechaCreacion)
		if err != nil {
			log.Fatal(err)
			return sliceProductos
		}
		sliceProductos = append(sliceProductos, productoLeido)
	}
	return sliceProductos

}
