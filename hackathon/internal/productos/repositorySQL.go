package internal

import (
	"context"
	"errors"

	"github.com/extmatperez/meli_bootcamp2/tree/pescie_juan/hackathon/internal/models"
	"github.com/extmatperez/meli_bootcamp2/tree/pescie_juan/hackathon/pkg/db"
)

type RepositorySQL interface {
	GetAll() ([]models.Product, error)
	// GetFullData() ([]models.Product, error)
	GetById(id int) (models.Product, error)
	Store(models.Product) (models.Product, error)
	Update(models.Product) (models.Product, error)
	UpdateContext(context.Context, models.Product) (models.Product, error)
	// UpdateNombrePrecio(id int, nombre string, precio float64) (models.Product, error)
	Delete(id int) error
	DeleteAll() error
}

const (
	queryGetOne    = "SELECT id, description, price FROM products WHERE id=?"
	queryGetAll    = "SELECT id, description, price FROM products"
	queryStore     = "INSERT INTO products(id, description, price) VALUES(?,?,?)"
	queryUpdate    = "UPDATE productos SET description = ?, price = ? WHERE id= ?"
	queryDelete    = "DELETE from productos WHERE id = ?"
	queryDeleteAll = "DELETE FROM products"
)

type repositorySql struct{}

func NewRepositorySQL() RepositorySQL {
	return &repositorySql{}
}

func (r *repositorySql) Store(p models.Product) (models.Product, error) {
	db := db.StorageDB
	statement, err := db.Prepare(queryStore)
	if err != nil {
		return models.Product{}, err
	}
	result, err := statement.Exec(p.ID, p.Description, p.Price)
	if err != nil {
		return models.Product{}, err
	}
	defer statement.Close()
	idCreado, _ := result.LastInsertId()
	p.ID = int(idCreado)
	return p, nil

}

func (r *repositorySql) GetById(id int) (models.Product, error) {
	db := db.StorageDB
	var p models.Product
	rows, err := db.Query(queryGetOne, id)
	if err != nil {
		return models.Product{}, err
	}
	for rows.Next() {
		err = rows.Scan(&p.ID, &p.Description, &p.Price)
		if err != nil {
			return models.Product{}, err
		}

	}
	return p, nil

}

// func (r *repositorySql) GetFullData() ([]models.Product, error) {
// 	db := db.StorageDB
// 	var productos []models.Product
// 	rows, err := db.Query(queryGetFullData)
// 	if err != nil {
// 		return nil, err
// 	}
// 	for rows.Next() {
// 		var p models.Product
// 		err = rows.Scan(&p.ID, &p.Description, &p.Price)
// 		if err != nil {
// 			return nil, err
// 		}
// 		fmt.Println("PRODUCTO : ", p)
// 		productos = append(productos, p)
// 	}
// 	return productos, nil
// }

func (r *repositorySql) GetAll() ([]models.Product, error) {
	db := db.StorageDB
	var productosLeidos []models.Product
	rows, err := db.Query(queryGetAll)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var p models.Product
		err = rows.Scan(&p.ID, &p.Description, &p.Price)
		if err != nil {
			return nil, err
		}
		productosLeidos = append(productosLeidos, p)

	}
	return productosLeidos, nil

}
func (r *repositorySql) Update(p models.Product) (models.Product, error) {
	db := db.StorageDB
	statement, err := db.Prepare(queryUpdate)
	if err != nil {
		return models.Product{}, err
	}
	result, err := statement.Exec(&p.ID, &p.Description, &p.Price)
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
func (r *repositorySql) UpdateContext(ctx context.Context, p models.Product) (models.Product, error) {
	db := db.StorageDB
	statement, err := db.Prepare(queryUpdate)
	if err != nil {
		return models.Product{}, err
	}
	result, err := statement.ExecContext(ctx, &p.ID, &p.Description, &p.Price)
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
func (r *repositorySql) DeleteAll() error {
	db := db.StorageDB
	statement, err := db.Prepare(queryDeleteAll)
	if err != nil {
		return err
	}
	defer statement.Close()
	_, err = statement.Exec()
	if err != nil {
		return err
	}

	return nil

}
