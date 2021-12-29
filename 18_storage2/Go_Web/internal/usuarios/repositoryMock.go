package internal

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log"

	"github.com/extmatperez/meli_bootcamp2/tree/aponte_nicolas/18_storage2/Go_Web/internal/models"
)

type RepositorySQLMock interface {
	Store(usuario models.Usuario) (models.Usuario, error)
	GetOne(id int) models.Usuario
	Update(ctx context.Context, usuario models.Usuario) (models.Usuario, error)
	GetByName(nombre string) ([]models.Usuario, error)
	GetAll() ([]models.Usuario, error)
	Delete(id int) error
}

type repositorySQLMock struct {
	db *sql.DB
}

func NewRepositorySQLMock(db *sql.DB) RepositorySQLMock {
	return &repositorySQLMock{db}
}

func (r *repositorySQLMock) Store(usuario models.Usuario) (models.Usuario, error) {
	stmt, err := r.db.Prepare("INSERT INTO users(nombre, apellido, email, edad, altura, activo, fecha_creacion) values (?,?,?,?,?,?,?)")
	if err != nil {
		log.Fatal(err)
	}
	result, err := stmt.Exec(usuario.Nombre, usuario.Apellido, usuario.Email, usuario.Edad, usuario.Altura, usuario.Activo, usuario.FechaCreacion)
	if err != nil {
		return models.Usuario{}, err
	}

	idCreado, _ := result.LastInsertId()
	usuario.ID = int(idCreado)

	return usuario, nil
}

func (r *repositorySQLMock) GetOne(id int) models.Usuario {
	var user models.Usuario
	rows, err := r.db.Query("SELECT id, nombre,apellido, email, edad, altura, activo,fecha_creacion FROM users WHERE id = ?", id)

	if err != nil {
		log.Fatal(err)
		return user
	}

	for rows.Next() {
		err = rows.Scan(&user.ID, &user.Nombre, &user.Apellido, &user.Email, &user.Edad, &user.Altura, &user.Activo, &user.FechaCreacion)
		if err != nil {
			log.Fatal(err)
			return user
		}
	}
	return user
}

func (r *repositorySQLMock) Update(ctx context.Context, usuario models.Usuario) (models.Usuario, error) {

	// stmt, err := db.Prepare("UPDATE users SET nombre = ?, apellido = ?, edad = ? WHERE id = ?")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// defer stmt.Close()
	//query := "SELECT SLEEP(30) FROM DUAL"
	query := "UPDATE users SET nombre = ?, apellido = ?, edad = ? WHERE id = ?"
	fmt.Println(usuario.ID)
	result, err := r.db.ExecContext(ctx, query, usuario.Nombre, usuario.Apellido, usuario.Edad, usuario.ID)
	if err != nil {
		fmt.Println(err)
		return models.Usuario{}, err
	}
	filasActualizadas, _ := result.RowsAffected()
	if filasActualizadas == 0 {
		return models.Usuario{}, errors.New("no se encontro el usuario")
	}

	return usuario, nil
}

func (r *repositorySQLMock) GetByName(nombre string) ([]models.Usuario, error) {

	var users []models.Usuario
	rows, err := r.db.Query("SELECT id, nombre,apellido, email, edad, altura, activo,fecha_creacion FROM users WHERE nombre = ?", nombre)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	for rows.Next() {
		var user models.Usuario
		if err = rows.Scan(&user.ID, &user.Nombre, &user.Apellido, &user.Email, &user.Edad, &user.Altura, &user.Activo, &user.FechaCreacion); err != nil {
			log.Fatal(err)
			return users, err
		}
		users = append(users, user)
	}

	return users, nil
}

func (r *repositorySQLMock) GetAll() ([]models.Usuario, error) {

	var users []models.Usuario
	rows, err := r.db.Query("SELECT id, nombre,apellido, email, edad, altura, activo,fecha_creacion FROM users")
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	for rows.Next() {
		var user models.Usuario
		if err = rows.Scan(&user.ID, &user.Nombre, &user.Apellido, &user.Email, &user.Edad, &user.Altura, &user.Activo, &user.FechaCreacion); err != nil {
			log.Fatal(err)
			return users, err
		}
		users = append(users, user)
	}

	return users, nil
}

func (r *repositorySQLMock) Delete(id int) error {

	stamt, err := r.db.Prepare("DELETE FROM users WHERE id = ?")
	if err != nil {
		log.Fatal(err)
	}
	defer stamt.Close()

	rows, err := stamt.Exec(id)
	if err != nil {
		log.Fatal(err)
		return err
	}

	filasAfectadas, _ := rows.RowsAffected()
	if filasAfectadas == 0 {
		return errors.New("no se encontro el usuario")
	}

	return nil
}
