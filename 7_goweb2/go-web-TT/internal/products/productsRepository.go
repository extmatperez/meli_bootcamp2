package internal

type Product struct {
	ID            int    `json:"id"`
	Nombre        string `json:"nombre"`
	Color         string `json:"color"`
	Precio        int    `json:"precio"`
	Stock         string `json:"stock"`
	Codigo        string `json:"codigo"`
	Publicado     bool   `json:"publicado"`
	FechaCreacion string `json:"fecha_creacion"`
}

var products []Product
var lastID int

type Repository interface {
	GetAll() ([]Product, error)
	Store(id int, nombre, color string, precio int, stock, codigo string, publicado bool, fechaCreacion string) (Product, error)
	LastID() (int, error)
}
type repository struct{}

func NewRepository() Repository {
	return &repository{}
}
func (repo *repository) GetAll() ([]Product, error) {
	return products, nil
}
func (repo *repository) Store(id int, nombre, color string, precio int, stock, codigo string, publicado bool, fechaCreacion string) (Product, error) {
	p := Product{id, nombre, color, precio, stock, codigo, publicado, fechaCreacion}
	lastID = id
	products = append(products, p)
	return p, nil
}
func (repo *repository) LastID() (int, error) {
	return lastID, nil
}
