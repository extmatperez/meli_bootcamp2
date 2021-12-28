package internal

import (
	"context"
	"database/sql"
	//"errors"
	"fmt"
	"log"

	"github.com/rossi_juancruz/meli_bootcamp2/8_goweb3/morning/ejercicios/internal/models"
	"github.com/rossi_juancruz/meli_bootcamp2/8_goweb3/morning/ejercicios/pkg/db"
)

const (
	QUERY_STORE = "INSERT INTO personas(nombre, apellido, edad) VALUES(?, ?, ?)"
	QUERY_GET_ONE = "SELECT id, nombre, apellido, edad, idciudad FROM personas WHERE id = ?"
	QUERY_GET_BY_NAME = "SELECT id, nombre, apellido, edad FROM personas WHERE nombre LIKE ?"
	QUERY_GET_ALL = "SELECT id, nombre, apellido, edad FROM personas"
	QUERY_UPDATE = "UPDATE personas SET nombre=?, apellido=?, edad=? WHERE id=?"
	//QUERY_UPDATE = "SELECT sleep(30) FROM dual"
	QUERY_DELETE = "DELETE FROM personas WHERE id=?"
	QUERY_GET_FULL_DATA_JOINED = "SELECT p.id, p.nombre, p.apellido, p.edad, c.id, c.nombreciudad, c.nombrepais FROM personas p INNER JOIN ciudad c ON p.idciudad = c.id"
)

type RepositorySQL interface {
	Store(persona models.Persona) (models.Persona, error)
	GetOne(id int) (models.Persona, error)
	GetOneWithContext(ctx context.Context, id int) (models.Persona, error)
	GetByName(nombre string) []models.Persona
	GetAll() ([]models.Persona, error)
	GetFullDataJoined() ([]models.Persona, error)
	UpdateWithContext(ctx context.Context, persona models.Persona) (models.Persona, error)
	Delete(id int) error	
}

type repositorySQL struct {}

var personasLeidas []models.Persona

func NewRepositorySQL() RepositorySQL {
	return &repositorySQL{}
}

func (r *repositorySQL) Store(persona models.Persona) (models.Persona, error) {
	db := db.StorageDB
	stmt, err := db.Prepare(QUERY_STORE)
	if err != nil {
		return models.Persona{}, err
	}

	defer stmt.Close()

	result, err := stmt.Exec(persona.Nombre, persona.Apellido, persona.Edad)
	if err != nil {
		return models.Persona{}, err
	}

	idCreado, _ := result.LastInsertId()
	persona.ID = int(idCreado)

	return persona, nil
}

func (r *repositorySQL) GetOne(id int) (models.Persona, error) {
	db := db.StorageDB
	personaLeida := models.Persona{}

	err := db.QueryRow(QUERY_GET_ONE, id).Scan(&personaLeida.ID, &personaLeida.Nombre, &personaLeida.Apellido, &personaLeida.Edad, &personaLeida.Domicilio.ID)
	if err != nil || err == sql.ErrNoRows {
		return models.Persona{}, err
	}

	return personaLeida, nil
}

func (r *repositorySQL) GetOneWithContext(ctx context.Context, id int) (models.Persona, error) {
	db := db.StorageDB
	personaLeida := models.Persona{}

	rows, err := db.QueryContext(ctx, QUERY_GET_ONE, id)
	if err != nil || err == sql.ErrNoRows {
		return models.Persona{}, err
	}

	for rows.Next() {
		if err := rows.Scan(&personaLeida.ID, &personaLeida.Nombre, &personaLeida.Apellido, &personaLeida.Edad, &personaLeida.Domicilio.ID); err != nil {
			return models.Persona{}, err
		}
	}

	return personaLeida, nil
}

func (r *repositorySQL) GetByName(nombre string) []models.Persona {
	db := db.StorageDB
	stmt, err := db.Prepare(QUERY_GET_BY_NAME)
    if err != nil {
        log.Fatal(err)
		return personasLeidas
    }
    rows, err := stmt.Query("%" + nombre + "%")
	if err != nil {
		log.Fatal(err)
		return personasLeidas
	}

	persona := models.Persona{}

	for rows.Next() {
		if err := rows.Scan(&persona.ID, &persona.Nombre, &persona.Apellido, &persona.Edad); err != nil {
			return personasLeidas
		}
		personasLeidas = append(personasLeidas, persona)
	}

	return personasLeidas
}

func (r *repositorySQL) GetAll() ([]models.Persona, error) {
	db := db.StorageDB
	var personas []models.Persona
	var personaLeida models.Persona

	rows, err := db.Query(QUERY_GET_ALL)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	
	for rows.Next() {
		err = rows.Scan(&personaLeida.ID, &personaLeida.Nombre, &personaLeida.Apellido, &personaLeida.Edad)
		if err != nil {
			log.Fatal(err)
			return nil, err
		}

		personas = append(personas, personaLeida)
	}

	return personas, nil
}

func (r *repositorySQL) GetFullDataJoined() ([]models.Persona, error) {
	db := db.StorageDB

	var personas []models.Persona
	var personaLeida models.Persona

	rows, err := db.Query(QUERY_GET_FULL_DATA_JOINED)

	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	for rows.Next() {
		err = rows.Scan(&personaLeida.ID, &personaLeida.Nombre, &personaLeida.Apellido, &personaLeida.Edad, &personaLeida.Domicilio.ID, &personaLeida.Domicilio.NombreCiudad, &personaLeida.Domicilio.NombrePais)
		if err != nil {
			log.Fatal(err)
			return nil, err
		}

		personas = append(personas, personaLeida)
	}

	return personas, nil
}

func (r *repositorySQL) UpdateWithContext(ctx context.Context, persona models.Persona) (models.Persona, error) {
	db := db.StorageDB
	stmt, err := db.Prepare(QUERY_UPDATE)
	if err != nil {
		return models.Persona{}, err
	}

	defer stmt.Close()

	res, err := stmt.ExecContext(ctx, &persona.Nombre, &persona.Apellido, &persona.Edad, &persona.ID)
	if err != nil {
		return models.Persona{}, err
	}

	num, err := res.RowsAffected()
	if err != nil || num == 0 {
		return models.Persona{}, err
	}

	return persona, nil
}

func (r *repositorySQL) Delete(id int) error {
	db := db.StorageDB
	stmt, err := db.Prepare(QUERY_DELETE)
	if err != nil {
		return err
	}

	defer stmt.Close()

	res, err := stmt.Exec(id)
	if err != nil {
		return err
	}

	affect, err := res.RowsAffected()
	if err != nil {
		return err
	}

	if affect < 1 {
		return fmt.Errorf("cannot delete, id: %v not found", id)
	}

	return nil
}