package internal

type Usuario struct {
	ID            int     `json:"id"`
	Nombre        string  `json:"nombre"`
	Apellido      string  `json:"apellido"`
	Email         string  `json:"email"`
	Edad          int     `json:"edad"`
	Altura        float64 `json:"altura"`
	Activo        bool    `json:"activo"`
	FechaCreacion string  `json:"fecha_creacion"`
}

var usuarios []Usuario
var lastID int

type Repository interface {
	GetAll() ([]Usuario, error)
	Store(id int, nombre string, apellido string, email string, edad int, altura float64, activo bool, fecha_creacion string) (Usuario, error)
	LastId() (int, error)
}

type repository struct{}

func NewRepository() Repository {
	return &repository{}
}

func (repo *repository) GetAll() ([]Usuario, error) {
	return usuarios, nil
}

func (repo *repository) Store(id int, nombre string, apellido string, email string, edad int, altura float64, activo bool, fecha_creacion string) (Usuario, error) {
	usr := Usuario{id, nombre, apellido, email, edad, altura, activo, fecha_creacion}
	lastID = id
	usuarios = append(usuarios, usr)
	return usr, nil
}

func (repo *repository) LastId() (int, error) {
	return lastID, nil
}
