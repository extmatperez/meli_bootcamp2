package internal

type Service interface {
	getAll() ([]Product, error)
	store(name string, color string, price float64, stock int64, code string, isPublished bool, createdAt string) (Product, error)
	loadFile() error
}

type service struct {
	repository Repository
}

func NewService(repository Repository) Service {
	return &service{repository: repository}
}

//servicio para obtener todos los productos
func (ser *service) getAll() ([]Product, error) {
	products, err := ser.repository.getAll()
	if err != nil {
		return nil, err
	}
	return products, nil
}

//Servicio para agregar un nuevo producto
func (ser *service) store(name string, color string, price float64, stock int64, code string, isPublished bool, createdAt string) (Product, error) {

	ultimoId := ser.repository.getLastID()

	product, err := ser.repository.store(ultimoId+1, name, color, price, stock, code, isPublished, createdAt)

	if err != nil {
		return Product{}, err
	}

	return product, nil
}

//Servicio para cargar el archivo de productos en memoria
func (ser *service) loadFile() error {
	err := ser.repository.loadFile()
	return err
}
