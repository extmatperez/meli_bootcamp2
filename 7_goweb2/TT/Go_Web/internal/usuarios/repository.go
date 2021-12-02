package internal

type Usuario struct {
	ID            int    `json:"id"`
	Nombre        string `json:"nombre"`
	Apellido      string `json:"apellido"`
	Email         string `json:"email"`
	Edad          int    `json:"edad"`
	Altura        int    `json:"altura"`
	Activo        bool   `json:"activo"`
	FechaCreacion string `json:"fecha_creacion"`
}

type Repository interface {
	GetAll() ([]Usuario, error)
	Store(id int, nombre, apellido, email string, edad, altura int, activo bool, fecha string) (Usuario, error)
	LastID() (int, error)
}

type repository struct{}

var usuarios []Usuario
var lastID int

func NewRepository() Repository {
	return &repository{}
}

func (repo *repository) GetAll() ([]Usuario, error) {
	return usuarios, nil
}

func (repo *repository) Store(id int, nombre, apellido, email string, edad, altura int, activo bool, fecha string) (Usuario, error) {
	lastID = id
	nuevoUsuario := Usuario{id, nombre, apellido, email, edad, altura, activo, fecha}
	usuarios = append(usuarios, nuevoUsuario)
	return nuevoUsuario, nil
}

func (repo *repository) LastID() (int, error) {
	return lastID, nil
}
