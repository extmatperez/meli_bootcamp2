package internal

import (
	"context"
	"errors"
	"fmt"

	"github.com/extmatperez/meli_bootcamp2/tree/pescie_juan/12_testing3/ejTM/internal/models"
	"github.com/extmatperez/meli_bootcamp2/tree/pescie_juan/12_testing3/ejTM/pkg/db"
)

type RepositorySQL interface {
	GetAll() ([]models.Producto, error)
	GetFullData() ([]models.Producto, error)
	GetById(id int) (models.Producto, error)
	Store(models.Producto) (models.Producto, error)
	Update(models.Producto) (models.Producto, error)
	UpdateContext(context.Context, models.Producto) (models.Producto, error)
	// UpdateNombrePrecio(id int, nombre string, precio float64) (models.Producto, error)
	Delete(id int) error
}

const (
	queryGetFullData = `SELECT p.id, p.nombre, p.color, p.precio, ciudad.nombre FROM productos as p INNER JOIN ciudades as ciudad
	ON p.id_ciudad = ciudad.id `
	queryGetOne = "SELECT id, nombre, color, precio FROM productos WHERE id=?"
	queryGetAll = "SELECT id, nombre, color, precio FROM productos"
	queryStore  = "INSERT INTO productos(nombre, color, precio) VALUES(?,?,?)"
	queryUpdate = "UPDATE productos SET nombre = ?, color = ?, precio = ? WHERE id= ?"
	queryDelete = "DELETE from productos WHERE id = ?"
)

type repositorySql struct{}

func NewRepositorySQL() RepositorySQL {
	return &repositorySql{}
}

func (r *repositorySql) Store(p models.Producto) (models.Producto, error) {
	db := db.StorageDB
	statement, err := db.Prepare(queryStore)
	if err != nil {
		return models.Producto{}, err
	}
	result, err := statement.Exec(p.Nombre, p.Color, p.Precio)
	if err != nil {
		return models.Producto{}, err
	}
	defer statement.Close()
	idCreado, _ := result.LastInsertId()
	p.Id = int(idCreado)
	return p, nil

}

func (r *repositorySql) GetById(id int) (models.Producto, error) {
	db := db.StorageDB
	var p models.Producto
	rows, err := db.Query(queryGetOne, id)
	if err != nil {
		return models.Producto{}, err
	}
	for rows.Next() {
		err = rows.Scan(&p.Id, &p.Nombre, &p.Color, &p.Precio)
		if err != nil {
			return models.Producto{}, err
		}

	}
	return p, nil

}
func (r *repositorySql) GetFullData() ([]models.Producto, error) {
	db := db.StorageDB
	var productos []models.Producto
	rows, err := db.Query(queryGetFullData)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var p models.Producto
		err = rows.Scan(&p.Id, &p.Nombre, &p.Color, &p.Precio, &p.Ciudad.Nombre)
		if err != nil {
			return nil, err
		}
		fmt.Println("PRODUCTO : ", p)
		productos = append(productos, p)
	}
	return productos, nil
}

func (r *repositorySql) GetAll() ([]models.Producto, error) {
	db := db.StorageDB
	var productosLeidos []models.Producto
	rows, err := db.Query(queryGetAll)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var p models.Producto
		err = rows.Scan(&p.Id, &p.Nombre, &p.Color, &p.Precio)
		if err != nil {
			return nil, err
		}
		productosLeidos = append(productosLeidos, p)

	}
	return productosLeidos, nil

}
func (r *repositorySql) Update(p models.Producto) (models.Producto, error) {
	db := db.StorageDB
	statement, err := db.Prepare(queryUpdate)
	if err != nil {
		return models.Producto{}, err
	}
	result, err := statement.Exec(p.Nombre, p.Color, p.Precio, p.Id)
	if err != nil {
		return models.Producto{}, err
	}
	filasActualizadas, _ := result.RowsAffected()
	if filasActualizadas == 0 {
		return models.Producto{}, errors.New("product not found")
	}
	defer statement.Close()

	return p, nil

}
func (r *repositorySql) UpdateContext(ctx context.Context, p models.Producto) (models.Producto, error) {
	db := db.StorageDB
	statement, err := db.Prepare(queryUpdate)
	if err != nil {
		return models.Producto{}, err
	}
	result, err := statement.ExecContext(ctx, p.Nombre, p.Color, p.Precio, p.Id)
	if err != nil {
		return models.Producto{}, err
	}
	filasActualizadas, _ := result.RowsAffected()
	if filasActualizadas == 0 {
		return models.Producto{}, errors.New("product not found")
	}
	defer statement.Close()

	return p, nil

}

func (r *repositorySql) Delete(id int) error {
	db := db.StorageDB
	statement, err := db.Prepare(queryDelete)
	if err != nil {
		return err
	}
	defer statement.Close()
	result, err := statement.Exec(id)
	if err != nil {
		return err
	}
	filasActualizadas, _ := result.RowsAffected()
	if filasActualizadas == 0 {
		return errors.New("product not found")
	}
	return nil

}
