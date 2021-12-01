package internal

type Service interface {
	GetAll() ([]Product, error)
	Store(name string, color string, price float64, stock int64, code string, isPublished bool, createdAt string) (Product, error)
	LoadFile() error
}

type service struct {
	repository Repository
}

func NewService(repository Repository) Service {
	return &service{repository: repository}
}

//servicio para obtener todos los productos
func (ser *service) GetAll() ([]Product, error) {
	products, err := ser.repository.GetAll()
	if err != nil {
		return nil, err
	}
	return products, nil
}

//Servicio para agregar un nuevo producto
func (ser *service) Store(name string, color string, price float64, stock int64, code string, isPublished bool, createdAt string) (Product, error) {

	ultimoId := ser.repository.GetLastID()

	product, err := ser.repository.Store(ultimoId+1, name, color, price, stock, code, isPublished, createdAt)

	if err != nil {
		return Product{}, err
	}

	return product, nil
}

//Servicio para cargar el archivo de productos en memoria
func (ser *service) LoadFile() error {
	err := ser.repository.LoadFile()
	return err
}
