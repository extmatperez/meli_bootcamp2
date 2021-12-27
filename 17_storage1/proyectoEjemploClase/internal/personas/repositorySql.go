package internal

import (
	"database/sql"
	"log"

	"github.com/extmatperez/meli_bootcamp2/tree/zamora_damian/17_storage1/proyectoEjemploClase/internal/models"
	"github.com/extmatperez/meli_bootcamp2/tree/zamora_damian/17_storage1/proyectoEjemploClase/pkg/db"
)

type RepositorySQL interface {
	Store(persona models.Persona) (models.Persona, error)
	GetOne(id int) models.Persona
	Update(persona models.Persona) (models.Persona, error)
	Delete(id int) error
	GetByName(nombre string) (models.Persona, error)
}

type repositorySQL struct{}

func NewRepositorioSQL() RepositorySQL {
	return &repositorySQL{}
}

func (r *repositorySQL) Store(persona models.Persona) (models.Persona, error) {
	//inicializamos la base de datos ( db.go) en /db
	db := db.StorageDB
	stmt, err := db.Prepare("INSERT INTO personas (nombre,apellido,edad) VALUES (?,?,?)")
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
	rows, err := db.Query("SELECT id,nombre,apellido,edad FROM personas WHERE id = ?", id)
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

func (r *repositorySQL) Update(persona models.Persona) (models.Persona, error) {
	db := db.StorageDB
	stmt, err := db.Prepare("UPDATE personas SET nombre = ?,apellido=?,edad=? WHERE id = ?")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()
	_, err = stmt.Exec(persona.Nombre, persona.Apellido, persona.Edad, persona.ID)
	if err != nil {
		return models.Persona{}, err
	}
	return persona, nil
}

func (r *repositorySQL) Delete(id int) error {
	db := db.StorageDB
	stmt, err := db.Prepare("DELETE FROM personas WHERE id = ?")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()
	_, err = stmt.Exec(id)
	if err != nil {
		return err
	}
	return nil
}

func (r *repositorySQL) GetByName(nombre string) (models.Persona, error) {
	db := db.StorageDB
	var persona models.Persona
	rows, err := db.Query("SELECT id,nombre,apellido,edad FROM personas WHERE nombre = ?", nombre)
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
