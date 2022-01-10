package internal

type Product struct {
	Id           int     `json:"id"`
	Name         string  `json:"name"`
	Color        string  `json:"color"`
	Price        float32 `json:"price"`
	Stock        int     `json:"stock"`
	Code         string  `json:"code"`
	Published    bool    `json:"published"`
	CreationDate string  `json:"creation_date"`
}

var products []Product

type Repository interface {
	GetAll() ([]Product, error)
	Save(newProduct Product) (Product, error)
}

type repository struct{}

func NewRepository() Repository {
	return &repository{}
}

func (repo *repository) GetAll() ([]Product, error) {
	return products, nil
}

func (repo *repository) Save(newProduct Product) (Product, error) {
	products = append(products, newProduct)
	return newProduct, nil
}
