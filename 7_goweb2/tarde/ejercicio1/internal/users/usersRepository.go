package internal

type User struct {
	ID            int    `json:"id"`
	Nombre        string `json:"nombre"`
	Apellido      string `json:"apellido"`
	Email         string `json:"email"`
	Edad          int    `json:"edad"`
	Altura        int    `json:"altura"`
	Activo        string `json:"activo"`
	FechaCreacion string `json:"fechaCreacion"`
}

var products []User
var lastID int

type Repository interface {
	GetAll() ([]User, error)
	Store(id int, nombre, apellido, email string, edad, altura int, activo, fechaCreacion string) (User, error)
	LastID() (int, error)
}
type repository struct{}

func NewRepository() Repository {
	return &repository{}
}
func (repo *repository) GetAll() ([]User, error) {
	return products, nil
}
func (repo *repository) Store(id int, nombre, apellido, email string, edad, altura int, activo, fechaCreacion string) (User, error) {
	p := User{id, nombre, apellido, email, edad, altura, activo, fechaCreacion}
	lastID = id
	products = append(products, p)
	return p, nil
}
func (repo *repository) LastID() (int, error) {
	return lastID, nil
}
