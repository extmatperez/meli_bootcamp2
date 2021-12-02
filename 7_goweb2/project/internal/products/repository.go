package internal

type Products struct {
	ID           int64  `json:"id"`
	Name         string `json:"name"`
	Color        string `json:"color"`
	Price        string `json:"price"`
	Stock        string `json:"stock"`
	Code         string `json:"code"`
	Published    bool   `json:"published"`
	CreationDate string `json:"creation_date"`
}

var products []Products
var lastID int64

type Repository interface {
	GetAll() ([]Products, error)
	Store(id int64, name string, color string, price string, stock string, code string, published bool, creationdate string) (Products, error)
	LastID() (int64, error)
}

type repository struct{}

func NewRepository() Repository {
	return &repository{}
}

func (r *repository) GetAll() ([]Products, error) {
	return products, nil
}

func (r *repository) Store(id int64, name string, color string, price string, stock string, code string, published bool, creationdate string) (Products, error) {
	prod := Products{id, name, color, price, stock, code, published, creationdate}
	products = append(products, prod)
	lastID = prod.ID
	return prod, nil
}

func (r *repository) LastID() (int64, error) {
	return lastID, nil
}
