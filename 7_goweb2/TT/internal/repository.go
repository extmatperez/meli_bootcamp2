package internal

type Product struct {
	ID    int     `json:"id"`
	Name  string  `json:"name"`
	Price float64 `json:"price"`
	Stock int     `json:"stock"`
}

var products []Product
var lastID int

type Repository interface {
	GetAll() ([]Product, error)
	Store(id int, name string, price float64, stock int) (Product, error)
	LastID() (int, error)
}

type repository struct{}

func NewRepository() Repository {
	return &repository{}
}

func (repo *repository) GetAll() ([]Product, error) {
	return products, nil
}

func (repo *repository) LastID() (int, error) {
	return lastID, nil
}

func (repo *repository) Store(id int, name string, price float64, stock int) (Product, error) {
	product := Product{id, name, price, stock}

	products = append(products, product)

	lastID = product.ID

	return product, nil
}
