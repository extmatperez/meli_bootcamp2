package internal

import (
	"errors"

	"github.com/extmatperez/meli_bootcamp2/tree/pescie_juan/12_testing3/ejTM/internal/models"
	"github.com/extmatperez/meli_bootcamp2/tree/pescie_juan/12_testing3/ejTM/pkg/db"
)

type RepositorySQL interface {
	GetAll() ([]models.Producto, error)
	GetById(id int) (models.Producto, error)
	Store(models.Producto) (models.Producto, error)
	// GetLastId() (int, error)
	Update(models.Producto) (models.Producto, error)
	// UpdateNombrePrecio(id int, nombre string, precio float64) (models.Producto, error)
	Delete(id int) error
}

const (
	queryGetOne = "SELECT id, nombre, color, precio FROM productos WHERE id=?"
	queryGetAll = "SELECT id, nombre, color, precio FROM productos"
	queryStore  = "INSERT INTO productos(nombre, color, precio, ) VALUES(?,?,?)"
	queryUpdate = "UPDATE productos SET nombre= ?, color= ?, precio= ? WHERE id= ?"
	queryDelete = "DELETE from productos WHERE id = ?"
)

type repositorySql struct{}

func NewRepositorySQL() RepositorySQL {
	return &repositorySql{}
}

func (r *repositorySql) Store(p models.Producto) (models.Producto, error) {
	db := db.StorageDB
	statement, err := db.Prepare("INSERT INTO productos(nombre, color, precio, ) VALUES(?,?,?)")
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
	query := "SELECT id, nombre, color, precio FROM productos WHERE id=?"
	var p models.Producto
	rows, err := db.Query(query)
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

func (r *repositorySql) GetAll() ([]models.Producto, error) {
	db := db.StorageDB
	var productosLeidos []models.Producto
	rows, err := db.Query("SELECT id, nombre, color, precio FROM productos")
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
	statement, err := db.Prepare("UPDATE productos SET nombre= ?, color= ?, precio= ? WHERE id= ?")
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

func (r *repositorySql) Delete(id int) error {
	db := db.StorageDB
	statement, err := db.Prepare()
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
