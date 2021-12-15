package internal

type Persona struct {
	Id       int    `json:"id"`
	Nombre   string `json:"nombre"`
	Apellido string `json:"apellido"`
	Email    string `json:"email"`
}

var personas []Persona
var lastId int

type Repository interface {
	GetAll() ([]Persona, error)
	Store(id int, nombre string, apellido string, email string) (Persona, error)
	Update(id int, nombre string, apellido string, email string) (Persona, error)
	LastId() (int, error)
}

//vamos a tener solo esta estructura
type repository struct{}

func NewRepository() Repository { /* retorna la interfaz , ya que tiene los metodos.*/
	// podriamos tener diferentes bases de datos, una rel y otra no rel por ejemplo--> (type repositoryNoRel struct)
	// en este caso en particular tenemos una sola qwue va a ser la strut repository, que es donde vamos a guardar
	// vamos a guardar lainformacion
	return &repository{}
	// &repository{} nos va a solicitar que implementemos lkso metodos de la interfaz

}

// implementamos metodo GetAll pasandole como parametros 1 nombre de variable, y 2 repository con puntero (seria la struct???)
// y va a devolver el un slice de []Persona y un error
func (repo *repository) GetAll() ([]Persona, error) {
	return personas, nil
}

func (repo *repository) Store(id int, nombre string, apellido string, email string) (Persona, error) {
	// creo que Store lo usamos como add. para agregar un objeto al slice
	per := Persona{id, nombre, apellido, email}
	// al id lo deberia calcular el service, PERO para este caso lo podriamos haber puesto aca ya que es autoincremental.!!! atencion
	lastId = id
	// le asignamos valor al id
	personas = append(personas, per)
	return per, nil
}

func (repo *repository) LastId() (int, error) {

	return lastId, nil

}

func (repo *repository) Update(id int, nombre string, apellido string, email string) (Persona, error) {
	return repo.Update(id, nombre, apellido, email)
}
