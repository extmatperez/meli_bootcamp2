package internal

import (
	"context"
	"database/sql"
	"errors"
	"log"

	"github.com/extmatperez/meli_bootcamp2/tree/de_bonis_matias/17_storage1/internal/models"
)

var (
	storeQuery       = "INSERT INTO products(name, type, count, price, id_warehouse) VALUES( ?, ?, ?, ?, ? )"
	getOneQuery      = "SELECT id, name, type, count, price FROM products WHERE id = ?"
	getAllQuery      = "SELECT id, name, type, count, price FROM products"
	updateQuery      = "UPDATE products SET name = ?, type = ?, count = ?, price = ? WHERE id = ?"
	deleteQuery      = "DELETE FROM products WHERE id = ?"
	getFullDataQuery = "select p.id, p.name, p.type, p.count, p.price, w.name, w.adress from products p inner join warehouses w on p.id_warehouse = w.id"
)

type RepositorySQL interface {
	Store(persona models.Producto) (models.Producto, error)
	GetOne(id int) models.Producto
	Update(context.Context, models.Producto) (models.Producto, error)
	GetAll() ([]models.Producto, error)
	Delete(id int) error
	GetFullData() ([]models.DTOProducto, error)

	//GetOneWithContext(ctx context.Context, id int) (models.Producto, error)
}

type repositorySQL struct {
	db *sql.DB
}

func NewRepositorySQL(dbModel *sql.DB) RepositorySQL {
	return &repositorySQL{db: dbModel}
}

func (r *repositorySQL) Store(producto models.Producto) (models.Producto, error) {
	stmt, err := r.db.Prepare(storeQuery)
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()
	var result sql.Result
	result, err = stmt.Exec(producto.Name, producto.Type, producto.Count, producto.Price, producto.WarehouseID)
	if err != nil {
		return models.Producto{}, err
	}
	idCreado, _ := result.LastInsertId()
	producto.ID = int(idCreado)

	return producto, nil
}

func (r *repositorySQL) GetOne(id int) models.Producto {
	var productoLeido models.Producto
	rows, err := r.db.Query(getOneQuery, id)

	if err != nil {
		log.Fatal(err)
		return productoLeido
	}

	for rows.Next() {
		err = rows.Scan(&productoLeido.ID, &productoLeido.Name, &productoLeido.Type, &productoLeido.Count, &productoLeido.Price)
		if err != nil {
			log.Fatal(err)
			return productoLeido
		}

	}
	return productoLeido
}
func (r *repositorySQL) GetAll() ([]models.Producto, error) {
	var misPersonas []models.Producto
	var productoLeido models.Producto
	rows, err := r.db.Query(getAllQuery)

	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	for rows.Next() {
		err = rows.Scan(&productoLeido.ID, &productoLeido.Name, &productoLeido.Type, &productoLeido.Count, &productoLeido.Price)
		if err != nil {
			log.Fatal(err)
			return nil, err
		}
		misPersonas = append(misPersonas, productoLeido)
	}
	return misPersonas, nil
}

func (r *repositorySQL) Update(ctx context.Context, producto models.Producto) (models.Producto, error) {
	stmt, err := r.db.Prepare(updateQuery)
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()
	result, err := stmt.ExecContext(ctx, producto.Name, producto.Type, producto.Count, producto.Price, producto.ID)
	if err != nil {
		return models.Producto{}, err
	}
	filasActualizadas, _ := result.RowsAffected()
	if filasActualizadas == 0 {
		return models.Producto{}, errors.New("no se encontro la persona")
	}

	return producto, nil
}

func (r *repositorySQL) Delete(id int) error {
	stmt, err := r.db.Prepare(deleteQuery)
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
		return errors.New("no se encontro la persona")
	}
	return nil
}

func (r *repositorySQL) GetFullData() ([]models.DTOProducto, error) {
	var misProductos []models.DTOProducto
	var productoLeido models.DTOProducto
	rows, err := r.db.Query(getFullDataQuery)

	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	for rows.Next() {
		err = rows.Scan(&productoLeido.ID, &productoLeido.Name, &productoLeido.Type, &productoLeido.Count, &productoLeido.Price, &productoLeido.Warehouse.Name, &productoLeido.Warehouse.Adress)
		if err != nil {
			log.Fatal(err)
			return nil, err
		}
		misProductos = append(misProductos, productoLeido)
	}
	return misProductos, nil
}

/* func (r *repositorySQL) GetOneWithContext(ctx context.Context, id int) (models.Producto, error) {
	db := db.StorageDB
	var productoLeido models.Producto
	// rows, err := db.QueryContext(ctx, "select sleep(30) from dual")
	rows, err := db.QueryContext(ctx, "SELECT id, nombre,apellido, edad FROM personas WHERE id = ?", id)

	if err != nil {
		log.Fatal(err)
		return productoLeido, err
	}

	for rows.Next() {
		err = rows.Scan(&productoLeido.ID, &productoLeido.Nombre, &productoLeido.Apellido, &productoLeido.Edad)
		if err != nil {
			log.Fatal(err)
			return productoLeido, err
		}

	}
	return productoLeido, nil
} */
