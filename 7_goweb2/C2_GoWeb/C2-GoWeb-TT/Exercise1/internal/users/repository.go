/*
Repositorio, debe tener el acceso a la variable guardada en memoria.
- Se debe crear el archivo repository.go
- Se debe crear la estructura de la entidad
- Se deben crear las variables globales donde guardar las entidades
- Se debe generar la interface Repository con todos sus métodos
- Se debe generar la estructura repository
- Se debe generar una función que devuelva el Repositorio
- Se deben implementar todos los métodos correspondientes a las operaciones a realizar (GetAll, Store, etc..)
*/
package internal

type User struct {
	ID          int    `json:"id"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	Email       string `json:"email"`
	Age         int    `json:"age"`
	Height      int    `json:"height"`
	Active      bool   `json:"active"`
	CrationDate string `json:"cration_date"`
}

var users []User
var lastID int

type Repository interface {
	GetAll() ([]User, error)
	Store(id int, first_name string, last_name string, email string, age int, height int, active bool, create_date string) (User, error)
	LastId() (int, error)
}

type repository struct{}

func NewRepository() Repository {
	return &repository{}
}

func (repo *repository) GetAll() ([]User, error) {
	return users, nil
}

func (repo *repository) Store(id int, first_name string, last_name string, email string, age int, height int, active bool, create_date string) (User, error) {
	us := User{id, first_name, last_name, email, age, height, active, create_date}
	users = append(users, us)
	lastID = us.ID
	return us, nil
}

func (repo *repository) LastId() (int, error) {
	return lastID, nil
}
