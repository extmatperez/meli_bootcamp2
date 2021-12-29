package internal

import (
	"context"
	"database/sql"
	"errors"
	"log"
	"strings"

	"github.com/extmatperez/meli_bootcamp2/17_storage1/TT/go-web/internal/models"
)

type RepositorySQL interface {
	GetByName(name string) (models.Product, error)
	Store(nombre string, color string, precio int, stock int, codigo string, publicado bool, fechaCreacion string) (models.Product, error)
	GetAll() ([]models.Product, error)
	Update(ctx context.Context, id int, nombre string, color string, precio int, stock int, codigo string, publicado bool, fechaCreacion string) (models.Product, error)
	UpdateNameAndPrice(id int, nombre string, precio int) (models.Product, error)
	Delete(id int) error
	GetOne(id int) (models.Product, error)
}

type repositorySQL struct {
	db *sql.DB
}

func NewRepositorySQL(db *sql.DB) RepositorySQL {
	return &repositorySQL{db}
}

func (r *repositorySQL) GetByName(name string) (models.Product, error) {

	var productReaded models.Product
	rows, err := r.db.Query("select id, nombre, color, precio, stock, codigo, publicado, fechaCreacion from products where LOWER(nombre) = ?", strings.ToLower(name))

	if err != nil {
		log.Fatal(err)
		return productReaded, err
	}

	for rows.Next() {
		err = rows.Scan(&productReaded.ID, &productReaded.Nombre, &productReaded.Color, &productReaded.Precio, &productReaded.Stock, &productReaded.Codigo, &productReaded.Publicado, &productReaded.FechaCreacion)
		if err != nil {
			log.Fatal(err)
			return productReaded, err
		}
	}
	return productReaded, nil
}

func (r *repositorySQL) Store(nombre string, color string, precio int, stock int, codigo string, publicado bool, fechaCreacion string) (models.Product, error) {

	stmt, err := r.db.Prepare("INSERT INTO products(nombre, color, precio, stock, codigo, publicado, fechaCreacion) VALUES( ?, ?, ? , ? , ? ,?, ?)")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()
	var result sql.Result
	result, err = stmt.Exec(nombre, color, precio, stock, codigo, publicado, fechaCreacion)
	if err != nil {
		return models.Product{}, err
	}
	idCreado, _ := result.LastInsertId()
	productCreated := models.Product{ID: int(idCreado), Nombre: nombre, Color: color, Precio: precio, Stock: stock, Codigo: codigo, Publicado: publicado, FechaCreacion: fechaCreacion}

	return productCreated, nil

}

func (r *repositorySQL) GetAll() ([]models.Product, error) {
	var productReaded models.Product
	prods := []models.Product{}
	rows, err := r.db.Query("select id, nombre, color, precio, stock, codigo, publicado, fechaCreacion from products ")

	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	for rows.Next() {
		err = rows.Scan(&productReaded.ID, &productReaded.Nombre, &productReaded.Color, &productReaded.Precio, &productReaded.Stock, &productReaded.Codigo, &productReaded.Publicado, &productReaded.FechaCreacion)
		if err != nil {
			log.Fatal(err)
			return nil, err
		}
		prods = append(prods, productReaded)
	}
	return prods, nil
}
func (r *repositorySQL) Update(ctx context.Context, id int, nombre string, color string, precio int, stock int, codigo string, publicado bool, fechaCreacion string) (models.Product, error) {

	stmt, err := r.db.Prepare("UPDATE products SET nombre = ?, color = ?, precio = ?, stock = ?, codigo = ?, publicado = ?, fechaCreacion = ? WHERE id = ?")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()
	result, err := stmt.ExecContext(ctx, nombre, color, precio, stock, codigo, publicado, fechaCreacion, id)
	if err != nil {
		return models.Product{}, err
	}
	filasActualizadas, _ := result.RowsAffected()
	if filasActualizadas == 0 {
		return models.Product{}, errors.New("No se encontro el producto")
	}
	productUpdated := models.Product{ID: id, Nombre: nombre, Color: color, Precio: precio, Stock: stock, Codigo: codigo, Publicado: publicado, FechaCreacion: fechaCreacion}

	return productUpdated, nil
}
func (r *repositorySQL) UpdateNameAndPrice(id int, nombre string, precio int) (models.Product, error) {

	stmt, err := r.db.Prepare("UPDATE products SET nombre = ?, precio = ?  WHERE id = ?")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()
	result, err := stmt.Exec(nombre, precio, id)
	if err != nil {
		return models.Product{}, err
	}
	filasActualizadas, _ := result.RowsAffected()
	if filasActualizadas == 0 {
		return models.Product{}, errors.New("No se encontro el producto")
	}
	productUpdated, err := r.GetOne(id)
	if err != nil {
		return models.Product{}, err
	}

	return productUpdated, nil
}

func (r *repositorySQL) Delete(id int) error {

	stmt, err := r.db.Prepare("DELETE FROM products WHERE id = ?")
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
		return errors.New("No se encontro el producto")
	}
	return nil
}

func (r *repositorySQL) GetOne(id int) (models.Product, error) {
	var productReaded models.Product
	row := r.db.QueryRow("select id, nombre, color, precio, stock, codigo, publicado, fechaCreacion from products where id= ?", id)
	err := row.Scan(&productReaded.ID, &productReaded.Nombre, &productReaded.Color, &productReaded.Precio, &productReaded.Stock, &productReaded.Codigo, &productReaded.Publicado, &productReaded.FechaCreacion)

	if err != nil {
		return models.Product{}, errors.New("Producto no encontrado")
	}

	return productReaded, nil
}
