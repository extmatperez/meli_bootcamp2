// Repository pertenece al paquete internal (carpeta general a la que pertenece)
package internal

import "fmt"

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
var last_id int

// Definimos métodos a utilizar en nuestro repo
type Repository interface {
	Get_users() ([]Users, error)
	Post_users(id int, first_name string, last_name string, email string, age int, height int, active bool, date string) (Users, error)
	Update_users(id int, first_name string, last_name string, email string, age int, height int, active bool, date string) (Users, error)
	Last_id() (int, error)
}

// Agregamos la estructura repository donde vamos a tener implementados los métodos de la interface
type repository struct{}

// Creo una función que retorne mi interface (si por ejemplo tengo otra interface para base de datos no relacionales agrego otra función que retorne esa interface)
func New_repository() Repository {
	return &repository{}
}

// Implementamos los métodos para que no marque erros la func new_repository
func (repo *repository) Get_users() ([]Users, error) {
	return users, nil
}
func (repo *repository) Post_users(id int, first_name string, last_name string, email string, age int, height int, active bool, date string) (Users, error) {
	user := Users{id, first_name, last_name, email, age, height, active, date}
	last_id = id
	users = append(users, user)
	return user, nil
}
func (repo *repository) Last_id() (int, error) {
	return last_id, nil
}
func (repo *repository) Update_users(id int, first_name string, last_name string, email string, age int, height int, active bool, date string) (Users, error) {
	user := Users{id, first_name, last_name, email, age, height, active, date}
	for i, v := range users {
		if v.ID == id {
			users[i] = user
			return user, nil
		}
	}
	return Users{}, fmt.Errorf("The user with id %v doesn't exist.", id)
}
