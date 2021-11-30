package internal

type Persona struct {
	Id       int    `json:"id"`
	Nombre   string `json:"nombre"`
	Apellido string `json:"apellido"`
	Edad     int    `json:"edad"`
}

var personas []Persona
var lastId int

type Repository interface {
	GetAll() ([]Persona, error)
	Store(id int, nombre string, apellido string, edad int) (Persona, error)
	//Store_bis(nuevaPersona Persona)(Persona, error)
	LastId() (int, error)
}

// Constructor vacio que solo va a tener implementados los metodos de la entidad.
type repository struct{}

func NewRepository() Repository {
	return &repository{}
}

// Implementacion de metodos en repository.
func (repo *repository) GetAll() ([]Persona, error) {
	return personas, nil
}

func (repo *repository) Store(id int, nombre string, apellido string, edad int) (Persona, error) {
	person := Persona{id, nombre, apellido, edad}
	lastId = id
	personas = append(personas, person)
	return person, nil
}

func (repo *repository) LastId() (int, error) {
	return lastId, nil
}
