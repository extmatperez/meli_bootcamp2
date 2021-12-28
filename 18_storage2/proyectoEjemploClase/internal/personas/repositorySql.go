package internal

import (
	"context"
	"database/sql"
	"errors"
	"log"

	"github.com/extmatperez/meli_bootcamp2/tree/zamora_damian/17_storage1/proyectoEjemploClase/internal/models"
	"github.com/extmatperez/meli_bootcamp2/tree/zamora_damian/17_storage1/proyectoEjemploClase/pkg/db"
)

type RepositorySQL interface {
	Store(persona models.Persona) (models.Persona, error)
	GetOne(id int) models.Persona
	Update(ctx context.Context, persona models.Persona) (models.Persona, error)
	Delete(id int) error
	GetByName(nombre string) (models.Persona, error)
	GetAll() ([]models.PersonaGetAllDTO, error)
}

type repositorySQL struct{}

func NewRepositorioSQL() RepositorySQL {
	return &repositorySQL{}
}

const (
	StoreQuery     = "INSERT INTO personas (nombre,apellido,edad) VALUES (?,?,?)"
	GetOneQuery    = "SELECT id,nombre,apellido,edad FROM personas WHERE id = ?"
	UpdateQuery    = "UPDATE personas SET nombre = ?,apellido=?,edad=? WHERE id = ?"
	DeleteQuery    = "DELETE FROM personas WHERE id = ?"
	GetByNameQuery = "SELECT id,nombre,apellido,edad FROM personas WHERE nombre = ?"
	GetAll         = "SELECT per.id, per.nombre,per.apellido,per.edad,ciu.ciudad,ciu.pais FROM personas as per INNER JOIN ciudad as ciu ON per.idciudad = ciu.id"
)

func (r *repositorySQL) Store(persona models.Persona) (models.Persona, error) {
	//inicializamos la base de datos ( db.go) en /db
	db := db.StorageDB
	stmt, err := db.Prepare(StoreQuery)
	if err != nil {
		log.Fatal(err)
	}
	//Se sentencia al terminar
	defer stmt.Close()
	var result sql.Result
	result, err = stmt.Exec(persona.Nombre, persona.Apellido, persona.Edad)
	if err != nil {
		return models.Persona{}, err
	}
	idInsertado, _ := result.LastInsertId()
	persona.ID = int(idInsertado)
	return persona, nil
}

func (r *repositorySQL) GetOne(id int) models.Persona {
	var persona models.Persona
	db := db.StorageDB
	rows, err := db.Query(GetOneQuery, id)
	if err != nil {
		log.Println(err)
		return persona
	}
	for rows.Next() {
		if err := rows.Scan(&persona.ID, &persona.Nombre, &persona.Apellido, &persona.Edad); err != nil {
			log.Println(err.Error())
			return persona
		}
	}
	return persona
}

func (r *repositorySQL) Update(ctx context.Context, persona models.Persona) (models.Persona, error) {
	db := db.StorageDB
	stmt, err := db.PrepareContext(context.Background(), UpdateQuery)
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()
	_, err = stmt.ExecContext(context.Background(), persona.Nombre, persona.Apellido, persona.Edad, persona.ID)
	if err != nil {
		return models.Persona{}, err
	}
	return persona, nil
}

func (r *repositorySQL) Delete(id int) error {
	db := db.StorageDB
	stmt, err := db.Prepare(DeleteQuery)
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
		return errors.New("No se ha eliminado")
	}

	return nil
}

func (r *repositorySQL) GetByName(nombre string) (models.Persona, error) {
	db := db.StorageDB
	var persona models.Persona
	rows, err := db.Query(GetByNameQuery, nombre)
	if err != nil {
		return models.Persona{}, err
	}
	for rows.Next() {
		if err := rows.Scan(&persona.ID, &persona.Nombre, &persona.Apellido, &persona.Edad); err != nil {
			log.Println(err.Error())
			return persona, err
		}
	}
	return persona, nil
}

func (r *repositorySQL) GetAll() ([]models.PersonaGetAllDTO, error) {
	var personas []models.PersonaGetAllDTO
	var persona models.PersonaGetAllDTO
	db := db.StorageDB
	rows, err := db.Query(GetAll)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	for rows.Next() {
		if err := rows.Scan(&persona.ID, &persona.Nombre, &persona.Apellido, &persona.Edad, &persona.Domicilio.Ciudad,
			&persona.Domicilio.Pais); err != nil {
			log.Println(err.Error())
			return nil, err
		}
		personas = append(personas, persona)
	}

	return personas, nil
}
