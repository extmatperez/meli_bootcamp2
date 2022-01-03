package internal

import (
	"context"
	"database/sql"
	"errors"
	"log"

	"github.com/extmatperez/meli_bootcamp2/tree/scerca_nahuel/19_storage3/ProyectoEstructura/internal/models"
	"github.com/extmatperez/meli_bootcamp2/tree/scerca_nahuel/19_storage3/ProyectoEstructura/pkg/db"
)

type RepositorySQLMock interface {
	GetByName(name string) ([]models.Producto, error)
	Store(producto models.Producto) (models.Producto, error)
	GetAll(ctx context.Context) ([]models.Producto, error)
	Update(ctx context.Context, producto models.Producto) (models.Producto, error)
}

type repositorySQLMock struct {
	db *sql.DB
}

func NewRepositorySQLMock(db *sql.DB) RepositorySQLMock {
	return &repositorySQLMock{db}
}

func (repsql *repositorySQLMock) Store(producto models.Producto) (models.Producto, error) {
	db := db.StorageDB

	stmt, err := db.Prepare("INSERT INTO productos(nombre,color,precio,stock,codigo,publicado,fechaCreacion) VALUES (?,?,?,?,?,?,?)")
	if err != nil {
		log.Fatal(err)
	}

	//desp de ejecutar la fucion esto se va a ejecutar
	defer stmt.Close()

	result, errExec := stmt.Exec(producto.Nombre, producto.Color, producto.Precio, producto.Stock, producto.Codigo, producto.Publicado, producto.FechaCreacion)
	if errExec != nil {
		return models.Producto{}, errExec
	}
	idCreado, _ := result.LastInsertId()
	producto.ID = int(idCreado)

	return producto, nil

}

func (r *repositorySQLMock) GetByName(name string) ([]models.Producto, error) {
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

func (r *repositorySQLMock) GetAll(ctx context.Context) ([]models.Producto, error) {
	var misProductos []models.Producto
	db := db.StorageDB
	var prodLeido models.Producto
	query := "SELECT id, nombre,color,precio,stock,codigo,publicado,fechaCreacion FROM productos"
	rows, err := db.QueryContext(ctx, query)

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
		misProductos = append(misProductos, prodLeido)
	}
	return misProductos, nil
}

func (repsql *repositorySQLMock) Update(ctx context.Context, producto models.Producto) (models.Producto, error) {
	db := db.StorageDB

	query := "UPDATE  productos SET nombre = ?,color= ?,precio= ?,stock= ?,codigo= ?,publicado= ?,fechaCreacion= ? WHERE id=?"
	stmt, err := db.Prepare(query)

	if err != nil {
		return models.Producto{}, err
	}
	defer stmt.Close()

	result, err := stmt.ExecContext(ctx, producto.Nombre, producto.Color, producto.Precio, producto.Stock, producto.Codigo, producto.Publicado, producto.FechaCreacion, producto.ID)

	if err != nil {
		return models.Producto{}, err
	}

	filasAfectadas, _ := result.RowsAffected()

	if filasAfectadas == 0 {
		return models.Producto{}, errors.New("Niguna fila afectada, no se encontro el producto")
	}
	return producto, nil
}
