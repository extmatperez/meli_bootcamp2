package internal

import (
	"database/sql"
	"log"

	"github.com/extmatperez/meli_bootcamp2/18_storage2/afternoon/go-web/internal/models"
)

const (
	insertProducto = `
		INSERT INTO productos (nombre, color, precio, stock, codigo, publicado, fechaCreacion)
		VALUES ( ?, ?, ?, ?, ?, ?, ? )
	`
	getOne = `
		SELECT id, nombre, color, precio, stock, codigo, publicado, fechaCreacion 
		FROM productos 
		WHERE id = ?
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
	deleteProducto = `
		DELETE FROM productos
		WHERE ID = ?;
	`
)

type RepositorySQLMock interface {
	Store(product models.Producto) (models.Producto, error)
	GetOne(id int) (models.Producto, error)
	Update(producto models.Producto, id int) (models.Producto, error)
	Delete(id int) error
}

type repositorySQLMock struct {
	db *sql.DB
}

func NewRepositorySQLMock(db *sql.DB) RepositorySQLMock {
	return &repositorySQLMock{db}
}

func (repo *repositorySQLMock) Store(producto models.Producto) (models.Producto, error) {
	// se prepara el SQL
	stmt, err := repo.db.Prepare(insertProducto)
	if err != nil {
		log.Fatal(err)
	}
	// se cierra la sentencia al terminar. Si quedan abiertas se genera consumos de memoria
	defer stmt.Close()
	var result sql.Result
	// del sql.Resul devuelto en la ejecución obtenemos el Id insertado
	result, err = stmt.Exec(producto.Nombre, producto.Color, producto.Precio, producto.Stock, producto.Codigo, producto.Publicado, producto.FechaCreacion) // retorna un sql.Result y un error
	if err != nil {
		return models.Producto{}, err
	}
	insertedId, _ := result.LastInsertId()
	producto.ID = int(insertedId)
	return producto, nil
}

func (repo *repositorySQLMock) GetOne(id int) (models.Producto, error) {

	var producto models.Producto
	rows, err := repo.db.Query(getOne, id)

	if err != nil {
		log.Println(err)
		return models.Producto{}, err
	}

	for rows.Next() {
		if err := rows.Scan(&producto.ID, &producto.Nombre, &producto.Color, &producto.Precio, &producto.Stock, &producto.Codigo, &producto.Publicado, &producto.FechaCreacion); err != nil {
			log.Println(err.Error())
			return models.Producto{}, err
		}
	}
	return producto, nil
}

func (repo *repositorySQLMock) Update(producto models.Producto, id int) (models.Producto, error) {
	// se prepara el SQL
	stmt, err := repo.db.Prepare(updateProducto)
	if err != nil {
		log.Fatal(err)
	}
	// se cierra la sentencia al terminar. Si quedan abiertas se genera consumos de memoria
	defer stmt.Close()

	// del sql.Resul devuelto en la ejecución obtenemos el Id insertado
	result, err := stmt.Exec(producto.Nombre, producto.Color, producto.Precio, producto.Stock, producto.Codigo, producto.Publicado, producto.FechaCreacion, id) // retorna un sql.Result y un error
	if err != nil {
		return models.Producto{}, err
	}
	updatedId, _ := result.LastInsertId()
	producto.ID = int(updatedId)
	return producto, nil
}

func (repo *repositorySQLMock) Delete(id int) error {

	stmt, err := repo.db.Prepare(deleteProducto)
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	result, err := stmt.Exec(id)
	if err != nil {
		return err
	}
	affected, _ := result.RowsAffected()
	if affected == 0 {
		return sql.ErrNoRows
	}
	return nil
}
