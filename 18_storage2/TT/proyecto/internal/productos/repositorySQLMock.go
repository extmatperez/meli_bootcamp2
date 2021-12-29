/*
Aplicar las buenas prácticas y recomendaciones para refactorizar el código de la capa repository.
	- Almacenar las queries como constantes.
	- No utilizar ”select *” en las queries.
*/

package internal

import (
	"context"
	"database/sql"
	"errors"

	"github.com/extmatperez/meli_bootcamp2/18_storage2/TT/proyecto/internal/models"

	_ "github.com/go-sql-driver/mysql"
)

type RepositorySQLMock interface {
	Store(nuevoProducto models.Producto) (models.Producto, error)
	GetAll() ([]models.Producto, error)
	Get(id int) (models.Producto, error)
	GetByName(nombre string) ([]models.Producto, error)
	Update(productoToUpdate models.Producto) (models.Producto, error)
	Delete(id int) error
	GetAllFullData() ([]models.Producto, error)
	GetWithContext(ctx context.Context, id int) (models.Producto, error)
	UpdateWithContext(ctx context.Context, productoToUpdate models.Producto) (models.Producto, error)
}

type repositorySQLMock struct {
	db *sql.DB
}

func NewRepositorySQLMock(db *sql.DB) RepositorySQLMock {
	return &repositorySQLMock{db}
}

func (r *repositorySQLMock) Store(nuevoProducto models.Producto) (models.Producto, error) {

	statement, err := r.db.Prepare("insert into productos(nombre, color, stock, precio, codigo, publicado, fecha_creacion, id_tipo) values(?,?,?,?,?,?,?,?);")

	if err != nil {
		return models.Producto{}, err
	}

	defer statement.Close()

	result, err := statement.Exec(nuevoProducto.Nombre, nuevoProducto.Color, nuevoProducto.Stock, nuevoProducto.Precio, nuevoProducto.Codigo, nuevoProducto.Publicado, nuevoProducto.FechaCreacion, nuevoProducto.Tipo.ID)

	if err != nil {
		return models.Producto{}, err
	}

	idCreado, _ := result.LastInsertId()

	nuevoProducto.ID = int(idCreado)

	return nuevoProducto, nil
}

func (r *repositorySQLMock) GetAll() ([]models.Producto, error) {

	rows, err := r.db.Query("select id, nombre, color, precio, stock, codigo, publicado, fecha_creacion from productos;")

	if err != nil {
		return nil, err
	}

	var productoLeido models.Producto
	var productos []models.Producto

	for rows.Next() {
		err := rows.Scan(&productoLeido.ID, &productoLeido.Nombre, &productoLeido.Color, &productoLeido.Precio, &productoLeido.Stock, &productoLeido.Codigo, &productoLeido.Publicado, &productoLeido.FechaCreacion)

		if err != nil {
			return nil, err
		}

		productos = append(productos, productoLeido)
	}

	return productos, nil
}

func (r *repositorySQLMock) Get(id int) (models.Producto, error) {

	rows, err := r.db.Query("select id, nombre, color, precio, stock, codigo, publicado, fecha_creacion from productos where id = ?;", id)

	if err != nil {
		return models.Producto{}, err
	}

	var productoLeido models.Producto

	if !rows.Next() {
		return models.Producto{}, errors.New("producto no encontrado")
	}

	for rows.Next() {
		err := rows.Scan(&productoLeido.ID, &productoLeido.Nombre, &productoLeido.Color, &productoLeido.Precio, &productoLeido.Stock, &productoLeido.Codigo, &productoLeido.Publicado, &productoLeido.FechaCreacion)

		if err != nil {
			return productoLeido, err
		}
	}

	return productoLeido, nil
}

func (r *repositorySQLMock) GetByName(nombre string) ([]models.Producto, error) {

	rows, err := r.db.Query("select id, nombre, color, precio, stock, codigo, publicado, fecha_creacion from productos where nombre like ?;", nombre)

	if err != nil {
		return nil, err
	}

	var productoLeido models.Producto
	var productos []models.Producto

	for rows.Next() {
		err := rows.Scan(&productoLeido.ID, &productoLeido.Nombre, &productoLeido.Color, &productoLeido.Precio, &productoLeido.Stock, &productoLeido.Codigo, &productoLeido.Publicado, &productoLeido.FechaCreacion)

		if err != nil {
			return nil, err
		}

		productos = append(productos, productoLeido)
	}

	return productos, nil
}

func (r *repositorySQLMock) Update(productoToUpdate models.Producto) (models.Producto, error) {

	statement, err := r.db.Prepare("update productos set nombre = ?, color = ?, stock = ?, precio = ?, codigo = ?, publicado = ?, fecha_creacion = ? where id = ?;")

	if err != nil {
		return models.Producto{}, err
	}

	defer statement.Close()

	var result sql.Result

	result, err = statement.Exec(productoToUpdate.Nombre, productoToUpdate.Color, productoToUpdate.Stock, productoToUpdate.Precio, productoToUpdate.Codigo, productoToUpdate.Publicado, productoToUpdate.FechaCreacion, productoToUpdate.ID)

	if err != nil {
		return models.Producto{}, err
	}

	rowsAffected, err := result.RowsAffected()

	if err != nil {
		return models.Producto{}, err
	}

	if rowsAffected == 0 {
		return models.Producto{}, errors.New("no se encontró el producto")
	}

	return productoToUpdate, nil
}

func (r *repositorySQLMock) Delete(id int) error {

	statement, err := r.db.Prepare("delete from productos where id = ?;")

	if err != nil {
		return err
	}

	defer statement.Close()

	result, err := statement.Exec(id)

	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()

	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return errors.New("no se encontró el producto")
	}

	return nil
}

func (r *repositorySQLMock) GetAllFullData() ([]models.Producto, error) {

	rows, err := r.db.Query("select p.id, p.nombre, p.color, p.precio, p.stock, p.codigo, p.publicado, p.fecha_creacion, t.id, t.descripcion from productos p join tipos t on p.id_tipo = t.id;")

	if err != nil {
		return nil, err
	}

	var productoLeido models.Producto
	var productos []models.Producto

	for rows.Next() {
		err := rows.Scan(&productoLeido.ID, &productoLeido.Nombre, &productoLeido.Color, &productoLeido.Precio, &productoLeido.Stock, &productoLeido.Codigo, &productoLeido.Publicado, &productoLeido.FechaCreacion, &productoLeido.Tipo.ID, &productoLeido.Tipo.Descripcion)

		if err != nil {
			return nil, err
		}

		productos = append(productos, productoLeido)
	}

	return productos, nil
}

func (r *repositorySQLMock) GetWithContext(ctx context.Context, id int) (models.Producto, error) {

	rows, err := r.db.QueryContext(ctx, "select id, nombre, color, precio, stock, codigo, publicado, fecha_creacion from productos where id = ?;", id)

	if err != nil {
		return models.Producto{}, err
	}

	var productoLeido models.Producto

	for rows.Next() {
		err := rows.Scan(&productoLeido.ID, &productoLeido.Nombre, &productoLeido.Color, &productoLeido.Precio, &productoLeido.Stock, &productoLeido.Codigo, &productoLeido.Publicado, &productoLeido.FechaCreacion)

		if err != nil {
			return productoLeido, err
		}
	}

	return productoLeido, nil
}

func (r *repositorySQLMock) UpdateWithContext(ctx context.Context, productoToUpdate models.Producto) (models.Producto, error) {

	statement, err := r.db.PrepareContext(ctx, "update productos set nombre = ?, color = ?, stock = ?, precio = ?, codigo = ?, publicado = ?, fecha_creacion = ? where id = ?;")

	if err != nil {
		return models.Producto{}, err
	}

	defer statement.Close()

	var result sql.Result

	result, err = statement.ExecContext(ctx, productoToUpdate.Nombre, productoToUpdate.Color, productoToUpdate.Stock, productoToUpdate.Precio, productoToUpdate.Codigo, productoToUpdate.Publicado, productoToUpdate.FechaCreacion, productoToUpdate.ID)

	if err != nil {
		return models.Producto{}, err
	}

	rowsAffected, err := result.RowsAffected()

	if err != nil {
		return models.Producto{}, err
	}

	if rowsAffected == 0 {
		return models.Producto{}, errors.New("no se encontró el producto")
	}

	return productoToUpdate, nil
}
