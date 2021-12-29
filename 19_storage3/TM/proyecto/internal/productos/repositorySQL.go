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
	"log"

	"github.com/extmatperez/meli_bootcamp2/19_storage3/TM/proyecto/internal/models"
	"github.com/extmatperez/meli_bootcamp2/19_storage3/TM/proyecto/pkg/db"

	_ "github.com/go-sql-driver/mysql"
)

type RepositorySQL interface {
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

type repositorySQL struct{}

func NewRepositorySQL() RepositorySQL {
	return &repositorySQL{}
}

func (r *repositorySQL) Store(nuevoProducto models.Producto) (models.Producto, error) {

	db := db.StorageDB
	statement, err := db.Prepare("insert into productos(nombre, color, stock, precio, codigo, publicado, fecha_creacion, id_tipo) values(?,?,?,?,?,?,?,?);")

	if err != nil {
		log.Fatal(err)
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

func (r *repositorySQL) GetAll() ([]models.Producto, error) {

	db := db.StorageDB
	rows, err := db.Query("select id, nombre, color, precio, stock, codigo, publicado, fecha_creacion from productos;")

	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	var productoLeido models.Producto
	var productos []models.Producto

	for rows.Next() {
		err := rows.Scan(&productoLeido.ID, &productoLeido.Nombre, &productoLeido.Color, &productoLeido.Precio, &productoLeido.Stock, &productoLeido.Codigo, &productoLeido.Publicado, &productoLeido.FechaCreacion)

		if err != nil {
			log.Fatal(err)
			return nil, err
		}

		productos = append(productos, productoLeido)
	}

	return productos, nil
}

func (r *repositorySQL) Get(id int) (models.Producto, error) {

	db := db.StorageDB
	rows, err := db.Query("select id, nombre, color, precio, stock, codigo, publicado, fecha_creacion from productos where id = ?;", id)

	if err != nil {
		log.Fatal(err)
		return models.Producto{}, err
	}

	var productoLeido models.Producto

	for rows.Next() {
		err := rows.Scan(&productoLeido.ID, &productoLeido.Nombre, &productoLeido.Color, &productoLeido.Precio, &productoLeido.Stock, &productoLeido.Codigo, &productoLeido.Publicado, &productoLeido.FechaCreacion)

		if err != nil {
			log.Fatal(err)
			return productoLeido, err
		}
	}

	return productoLeido, nil
}

func (r *repositorySQL) GetByName(nombre string) ([]models.Producto, error) {

	db := db.StorageDB
	rows, err := db.Query("select id, nombre, color, precio, stock, codigo, publicado, fecha_creacion from productos where nombre like ?;", nombre)

	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	var productoLeido models.Producto
	var productos []models.Producto

	for rows.Next() {
		err := rows.Scan(&productoLeido.ID, &productoLeido.Nombre, &productoLeido.Color, &productoLeido.Precio, &productoLeido.Stock, &productoLeido.Codigo, &productoLeido.Publicado, &productoLeido.FechaCreacion)

		if err != nil {
			log.Fatal(err)
			return nil, err
		}

		productos = append(productos, productoLeido)
	}

	return productos, nil
}

func (r *repositorySQL) Update(productoToUpdate models.Producto) (models.Producto, error) {

	db := db.StorageDB
	statement, err := db.Prepare("update productos set nombre = ?, color = ?, stock = ?, precio = ?, codigo = ?, publicado = ?, fecha_creacion = ? where id = ?;")

	if err != nil {
		log.Fatal(err)
	}

	defer statement.Close()

	var result sql.Result

	result, err = statement.Exec(productoToUpdate.Nombre, productoToUpdate.Color, productoToUpdate.Stock, productoToUpdate.Precio, productoToUpdate.Codigo, productoToUpdate.Publicado, productoToUpdate.FechaCreacion, productoToUpdate.ID)

	if err != nil {
		return models.Producto{}, err
	}

	rowsAffected, err := result.RowsAffected()

	if err != nil {
		log.Fatal(err)
	}

	if rowsAffected == 0 {
		return models.Producto{}, errors.New("no se encontró el producto")
	}

	return productoToUpdate, nil
}

func (r *repositorySQL) Delete(id int) error {

	db := db.StorageDB
	statement, err := db.Prepare("delete from productos where id = ?;")

	if err != nil {
		log.Fatal(err)
	}

	defer statement.Close()

	result, err := statement.Exec(id)

	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()

	if err != nil {
		log.Fatal(err)
	}

	if rowsAffected == 0 {
		return errors.New("no se encontró el producto")
	}

	return nil
}

func (r *repositorySQL) GetAllFullData() ([]models.Producto, error) {

	db := db.StorageDB
	rows, err := db.Query("select p.id, p.nombre, p.color, p.precio, p.stock, p.codigo, p.publicado, p.fecha_creacion, t.id, t.descripcion from productos p join tipos t on p.id_tipo = t.id;")

	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	var productoLeido models.Producto
	var productos []models.Producto

	for rows.Next() {
		err := rows.Scan(&productoLeido.ID, &productoLeido.Nombre, &productoLeido.Color, &productoLeido.Precio, &productoLeido.Stock, &productoLeido.Codigo, &productoLeido.Publicado, &productoLeido.FechaCreacion, &productoLeido.Tipo.ID, &productoLeido.Tipo.Descripcion)

		if err != nil {
			log.Fatal(err)
			return nil, err
		}

		productos = append(productos, productoLeido)
	}

	return productos, nil
}

func (r *repositorySQL) GetWithContext(ctx context.Context, id int) (models.Producto, error) {

	db := db.StorageDB
	rows, err := db.QueryContext(ctx, "select id, nombre, color, precio, stock, codigo, publicado, fecha_creacion from productos where id = ?;", id)

	if err != nil {
		log.Fatal(err)
		return models.Producto{}, err
	}

	var productoLeido models.Producto

	for rows.Next() {
		err := rows.Scan(&productoLeido.ID, &productoLeido.Nombre, &productoLeido.Color, &productoLeido.Precio, &productoLeido.Stock, &productoLeido.Codigo, &productoLeido.Publicado, &productoLeido.FechaCreacion)

		if err != nil {
			log.Fatal(err)
			return productoLeido, err
		}
	}

	return productoLeido, nil
}

func (r *repositorySQL) UpdateWithContext(ctx context.Context, productoToUpdate models.Producto) (models.Producto, error) {

	db := db.StorageDB
	statement, err := db.PrepareContext(ctx, "update productos set nombre = ?, color = ?, stock = ?, precio = ?, codigo = ?, publicado = ?, fecha_creacion = ? where id = ?;")

	if err != nil {
		log.Fatal(err)
	}

	defer statement.Close()

	var result sql.Result

	result, err = statement.ExecContext(ctx, productoToUpdate.Nombre, productoToUpdate.Color, productoToUpdate.Stock, productoToUpdate.Precio, productoToUpdate.Codigo, productoToUpdate.Publicado, productoToUpdate.FechaCreacion, productoToUpdate.ID)

	if err != nil {
		return models.Producto{}, err
	}

	rowsAffected, err := result.RowsAffected()

	if err != nil {
		log.Fatal(err)
	}

	if rowsAffected == 0 {
		return models.Producto{}, errors.New("no se encontró el producto")
	}

	return productoToUpdate, nil
}