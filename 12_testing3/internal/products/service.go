package internal

type service struct {
	repository Repository
}

type Service interface {
	GetAll() ([]Product, error)
	AddProduct(name, color string, price float64, stock, code int, published string, created string) (Product, error)
	UpdateProduct(id int, name, color string, price float64, stock, code int, published string, created string) (Product, error)
}

func NewService(repository Repository) Service {
	return &service{repository: repository}
}

func (serv *service) GetAll() ([]Product, error) {
	prod, err := serv.repository.GetAll()
	if err != nil {
		return nil, err
	}
	return prod, nil
	///// el control de errores deberia estar en el repo, entonces solo hace falta: ////////
	// return serv.repository.getAll()  //// y retorna lo que haya retornado el repo
}

func (serv *service) AddProduct(name, color string, price float64, stock, code int, published string, created string) (Product, error) {
	lastID := lastIDrepo + 1

	newProd, err := serv.repository.AddProduct(lastID, name, color, price, stock, code, published, created)

	if err != nil {
		return Product{}, err
	}
	return newProd, nil

}

func (serv *service) UpdateProduct(id int, name, color string, price float64, stock, code int, published string, created string) (Product, error) {
	return serv.repository.UpdateProduct(id, name, color, price, stock, code, published, created)

}
