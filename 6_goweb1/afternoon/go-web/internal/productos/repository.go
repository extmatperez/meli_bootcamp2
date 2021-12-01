package internal

type Producto struct {
	Id            int     `json:"id"`
	Nombre        string  `json:"nombre"`
	Color         string  `json:"color"`
	Precio        float64 `json:"precio"`
	Stock         int     `json:"stock"`
	Codigo        string  `json:"codigo"`
	Publicado     bool    `json:"publicado"`
	FechaCreacion string  `json:"fechaCreacion"`
}

var productos []Producto
var lastID int

type Repository interface {
	GetAll() ([]Producto, error)
	Store(id int, nombre string, color string, precio float64, stock int, codigo string, publicado bool, fechaCreacion string) (Producto, error)
	LastId() (int, error)
}

type repository struct{}

func NewRepository() Repository {
	return &repository{}
}

func (repo *repository) GetAll() ([]Producto, error) {
	return productos, nil
}

func (repo *repository) Store(id int, nombre string, color string, precio float64, stock int, codigo string, publicado bool, fechaCreacion string) (Producto, error) {
	per := Producto{id, nombre, color, precio, stock, codigo, publicado, fechaCreacion}
	lastID = id
	productos = append(productos, per)
	return per, nil
}

func (repo *repository) LastId() (int, error) {
	return lastID, nil
}
