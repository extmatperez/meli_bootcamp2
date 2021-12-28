package internal

import (
	"errors"

	"github.com/extmatperez/meli_bootcamp2/tree/arevalo_ivan/17_storage1/go_web/internal/models"
	"github.com/extmatperez/meli_bootcamp2/tree/arevalo_ivan/17_storage1/go_web/pkg/store"
)

type RepositorySQL interface {
	GetAll() ([]models.Product, error)
	GetById(id int) (models.Product, error)
	Store(models.Product) (models.Product, error)
	Update(models.Product) (models.Product, error)
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

func (r *repositorySql) Store(p models.Product) (models.Product, error) {
	db := store.StorageDB
	statement, err := db.Prepare(queryStore)
	if err != nil {
		return models.Product{}, err
	}
	result, err := statement.Exec(p.Name, p.Type, p.Count, p.Price)
	if err != nil {
		return models.Product{}, err
	}
	defer statement.Close()
	idCreado, _ := result.LastInsertId()
	p.Id = int(idCreado)
	return p, nil

}

func (r *repositorySql) GetById(id int) (models.Product, error) {
	db := store.StorageDB
	var p models.Product
	rows, err := db.Query(queryGetOne)
	if err != nil {
		return models.Product{}, err
	}
	for rows.Next() {
		err = rows.Scan(&p.Id, &p.Name, &p.Type, &p.Count, &p.Price)
		if err != nil {
			return models.Product{}, err
		}

	}
	return p, nil

}

func (r *repositorySql) GetAll() ([]models.Product, error) {
	db := store.StorageDB
	var productosLeidos []models.Product
	rows, err := db.Query(queryGetAll)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var p models.Product
		err = rows.Scan(&p.Id, &p.Name, &p.Type, &p.Count, &p.Price)
		if err != nil {
			return nil, err
		}
		productosLeidos = append(productosLeidos, p)

	}
	return productosLeidos, nil

}
func (r *repositorySql) Update(p models.Product) (models.Product, error) {
	db := store.StorageDB
	statement, err := db.Prepare(queryUpdate)
	if err != nil {
		return models.Product{}, err
	}
	result, err := statement.Exec(p.Name, p.Type, p.Count, p.Price)
	if err != nil {
		return models.Product{}, err
	}
	filasActualizadas, _ := result.RowsAffected()
	if filasActualizadas == 0 {
		return models.Product{}, errors.New("product not found")
	}
	defer statement.Close()

	return p, nil

}

func (r *repositorySql) Delete(id int) error {
	db := store.StorageDB
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
