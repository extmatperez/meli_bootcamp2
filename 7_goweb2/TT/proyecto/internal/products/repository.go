package internal

type Product struct {
	ID              int     `json:"id"`
	Nombre          string  `json:"nombre"`
	Color           string  `json:"color"`
	Precio          float64 `json:"precio"`
	Stock           int     `json:"stock"`
	Codigo          string  `json:"codigo"`
	Publicado       bool    `json:"publicado"`
	FechaDeCreacion string  `json:"fecha_de_creacion"`
}

var products []Product
var lastID int

type Repository interface {
	GetAll() ([]Product, error)
	Store(id int, nombre, color string, precio float64, stock int, codigo string, publicado bool, fechaDeCreacion string) (Product, error)
	//Store2(nuevoProduct Product)(Product,err error)
	LastId() (int, error)
}

type repository struct{}

func NewRepository() Repository {
	return &repository{}
}

func (repo *repository) GetAll() ([]Product, error) {
	return products, nil
}

func (repo *repository) Store(id int, nombre, color string, precio float64, stock int, codigo string, publicado bool, fechaDeCreacion string) (Product, error) {
	prod := Product{id, nombre, color, precio, stock, codigo, publicado, fechaDeCreacion}
	lastID = id
	products = append(products, prod)
	return prod, nil
}

func (repo *repository) LastId() (int, error) {
	return lastID, nil
}
