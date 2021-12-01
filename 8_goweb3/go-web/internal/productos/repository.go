package internal

type Producto struct {
	ID        int    `json:"id"`
	Nombre    string `json:"nombre"`
	Color     string `json:"color"`
	Precio    string `json:"precio"`
	Stock     int    `json:"stock"`
	Codigo    string `json:"codigo"`
	Publicado bool   `json:"publicado"`
	Creado    string `json:"creado"`
}

var ps []Producto
var lastID int

type Repository interface {
	GetAll() ([]Producto, error)
	Store(id int, nombre, color, precio string, stock int, codigo string, publicado bool, creado string) (Producto, error)
	LastID() (int, error)
}

type repository struct{}

func NewRepository() Repository {
	return &repository{}
}

func (r *repository) GetAll() ([]Producto, error) {
	return ps, nil
}

func (r *repository) LastID() (int, error) {
	return lastID, nil
}

func (r *repository) Store(id int, nombre, color, precio string, stock int, codigo string, publicado bool, creado string) (Producto, error) {
	p := Producto{id, nombre, color, precio, stock, codigo, publicado, creado}
	ps = append(ps, p)
	lastID = p.ID
	return p, nil
}
