// Repository pertenece al paquete internal (carpeta general a la que pertenece)
package internal

import (
	"fmt"

	"github.com/extmatperez/meli_bootcamp2/tree/montenegro_edgar/9_goweb4/morning_activities/Exercise_1/pkg/store"
)

// Estructura de los datos que voy a manipular
type Users struct {
	ID        int    `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	Age       int    `json:"age"`
	Height    int    `json:"height"`
	Active    bool   `json:"active"`
	Date      string `json:"date"`
}

// Mientras no tengamos base de datos necesitamos saber donde almacenar los datos
var users []Users

// Definimos métodos a utilizar en nuestro repo
type Repository interface {
	Get_users() ([]Users, error)
	Post_users(id int, first_name string, last_name string, email string, age int, height int, active bool, date string) (Users, error)
	Update_users(id int, first_name string, last_name string, email string, age int, height int, active bool, date string) (Users, error)
	Update_users_first_name(id int, first_name string) (Users, error)
	Delete_users(id int) error
	Last_id() (int, error)
}

// Agregamos la estructura repository donde vamos a tener implementados los métodos de la interface
// * Dentro de la estructura repository se declara el campo de tipo Store que se importará del paquete que se generó previamente.
type repository struct {
	db store.Store
}

// Creo una función que retorne mi interface (si por ejemplo tengo otra interface para base de datos no relacionales agrego otra función que retorne esa interface)
//* Agregamos como argumento el store.
func New_repository(db store.Store) Repository {
	return &repository{db}
}

// Implementamos los métodos para que no marque erros la func new_repository
// Implementamos la funcionalidad para llamar a todos los usuarios existentes en el slice
func (repo *repository) Get_users() ([]Users, error) {
	err := repo.db.Read(&users)
	if err != nil {
		return nil, err
	}
	return users, nil
}

// Implementamos la funcionalidad para agregar usuarios a nuestro slice
func (repo *repository) Post_users(id int, first_name string, last_name string, email string, age int, height int, active bool, date string) (Users, error) {
	// Leemos lo que exista en nuestro slice (base de datos)
	err := repo.db.Read(&users)
	if err != nil {
		return Users{}, err
	}

	user := Users{id, first_name, last_name, email, age, height, active, date}
	users = append(users, user)
	err = repo.db.Write(users)

	if err != nil {
		return Users{}, err
	}
	return user, nil
}

// Implementamos la funcionalidad para obtener el ID del último usuario en el slice
func (repo *repository) Last_id() (int, error) {

	err := repo.db.Read(&users)
	if err != nil {
		return 0, err
	}
	if len(users) == 0 {
		return 0, nil
	}
	return users[len(users)-1].ID, nil
}

// Implementamos la funcionalidad para actualizar el usuario en memoria, en caso de que coincida con el ID enviado, caso contrario retorna un error.
func (repo *repository) Update_users(id int, first_name string, last_name string, email string, age int, height int, active bool, date string) (Users, error) {
	// Leemos para que levante todo lo que hay en el JSON
	err := repo.db.Read(&users)
	if err != nil {
		return Users{}, err
	}
	user := Users{id, first_name, last_name, email, age, height, active, date}
	// Recorro el JSON
	for i, v := range users {
		if v.ID == id {
			// Actualizo mi JSON
			users[i] = user
			// Sobre-escribo mi JSON
			err := repo.db.Write(users)
			if err != nil {
				return Users{}, err
			}
			return user, nil
		}
	}
	return Users{}, fmt.Errorf("the user with id %v doesn't exist", id)
}

// Implementamos la funcionalidad para actualizar el usuario en memoria, en caso de que coincida con el ID enviado, caso contrario retorna un error.
func (repo *repository) Update_users_first_name(id int, first_name string) (Users, error) {
	err := repo.db.Read(&users)

	if err != nil {
		return Users{}, err
	}

	for i, v := range users {
		if v.ID == id {
			users[i].FirstName = first_name
			err := repo.db.Write(users)
			if err != nil {
				return Users{}, err
			}
			return users[i], nil
		}
	}
	return Users{}, fmt.Errorf("the user with id %v doesn't exist, try with another user to change the name", id)
}

// Implementamos la funcionalidad para borrar el usuario en memoria.
func (repo *repository) Delete_users(id int) error {
	// Otra manera de eliminar sería...
	/* var new_users []Users

	for _, v := range users {
		if v.ID != id {
			new_users = append(new_users, v)
		}
	}
	users = new_users */

	err := repo.db.Read(&users)
	if err != nil {
		return err
	}
	index := 0

	for i, v := range users {
		if v.ID == id {
			index = i
			users = append(users[:index], users[index+1:]...)
			err := repo.db.Write(users)
			return err
		}
	}
	return fmt.Errorf("the user with id %v doesn't exist", id)
}
