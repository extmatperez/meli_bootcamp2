package internal

import (
	"context"
	"database/sql"
	"log"

	"github.com/extmatperez/meli_bootcamp2/18_storage2/morning/go-web/internal/models"
	"github.com/extmatperez/meli_bootcamp2/18_storage2/morning/go-web/pkg/db"
)

type RepositorySQL interface {
	Store(product models.Producto) (models.Producto, error)
	GetByName(name string) ([]models.Producto, error)
	GetAll() ([]models.Producto, error)
	Update(ctx context.Context, producto models.Producto, id int) (models.Producto, error)
}

type repositorySQL struct{}

func NewRepositorySQL() RepositorySQL {
	return &repositorySQL{}
}

const (
	insertProducto = `
		INSERT INTO productos (nombre, color, precio, stock, codigo, publicado, fechaCreacion)
		VALUES ( ?, ?, ?, ?, ?, ?, ? )
	`
	getByName = `
		SELECT id, nombre, color, precio, stock, codigo, publicado, fechaCreacion 
		FROM productos 
		WHERE nombre = ?
	`
	getAll = `
		SELECT id, nombre, color, precio, stock, codigo, publicado, fechaCreacion 
		FROM productos 
	`
	/* updateProducto = `
		UPDATE productos (nombre, color, precio, stock, codigo, publicado, fechaCreacion)
		VALUES ( ?, ?, ?, ?, ?, ?, ? )
		WHERE id = ?
	` */
	updateProducto = `
		UPDATE productos
		SET
		Nombre = ?,
		Color = ?,
		Precio = ?,
		Stock = ?,
		Codigo = ?,
		Publicado = ?,
		FechaCreacion = ?
		WHERE ID = ?;
	`
)

func (repo *repositorySQL) Store(producto models.Producto) (models.Producto, error) {
	db := db.StorageDB                      // se inicializa la base
	stmt, err := db.Prepare(insertProducto) // se prepara el SQL
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

	rows, err := db.Query(getByName, name)
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

func (repo *repositorySQL) GetAll() ([]models.Producto, error) {
	var producto models.Producto
	var productos []models.Producto
	db := db.StorageDB

	rows, err := db.Query(getAll)
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

func (repo *repositorySQL) Update(ctx context.Context, producto models.Producto, id int) (models.Producto, error) {
	// se inicializa la base
	db := db.StorageDB
	// se prepara el SQL
	stmt, err := db.PrepareContext(ctx, updateProducto)
	if err != nil {
		log.Fatal(err)
	}
	// se cierra la sentencia al terminar. Si quedan abiertas se genera consumos de memoria
	defer stmt.Close()
	// retorna un sql.Result y un error
	_, err = stmt.ExecContext(ctx, producto.Nombre, producto.Color, producto.Precio, producto.Stock, producto.Codigo, producto.Publicado, producto.FechaCreacion, id)
	if err != nil {
		return models.Producto{}, err
	}

	producto.ID = id
	return producto, nil
}
