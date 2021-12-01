package internal

type Producto struct {
	ID            int     `json:"id"`
	Nombre        string  `json:"nombre"`
	Color         string  `json:"color"`
	Precio        float64 `json:"precio"`
	Stock         int     `json:"stock"`
	Codigo        string  `json:"codigo"`
	Publicado     bool    `json:"publicado"`
	FechaCreacion string  `json:"fecha_creacion"`
}

var productos []Producto
var lastID int

type RepositoryProducto interface {
	GetAll() ([]Producto, error)
	Store(id int, nombre string, color string, precio float64, stock int, codigo string, publicado bool, fechaCreacion string) (Producto, error)
	//Store2(nuevaPersona Persona)(Persona, error)
	LastId() (int, error)
}

type repositoryProducto struct{}

func NewRepository() RepositoryProducto {
	return &repositoryProducto{}
}

func (repo *repositoryProducto) GetAll() ([]Producto, error) {
	return productos, nil
}

func (repo *repositoryProducto) Store(id int, nombre string, color string, precio float64, stock int, codigo string, publicado bool, fechaCreacion string) (Producto, error) {
	prod := Producto{id, nombre, color, precio, stock, codigo, publicado, fechaCreacion}
	lastID = id
	productos = append(productos, prod)
	return prod, nil
}

func (repo *repositoryProducto) LastId() (int, error) {
	return lastID, nil
}
