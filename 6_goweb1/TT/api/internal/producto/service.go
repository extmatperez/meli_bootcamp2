package internal

type Service interface {
	GetAll() ([]Product, error)
	Store(name string, color string, price float64, stock int64, code string, isPublished bool, createdAt string) (Product, error)
	LoadFile() error
	Delete(id int64) (string, error)
	Update(id int64, name string, color string, price float64, stock int64, code string, isPublished bool, createdAt string) (Product, error)
	UpdateNombre(id int64, name string) (Product, error)
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

func (ser *service) Delete(id int64) (string, error) {
	mes, err := ser.repository.Delete(id)

	if err != nil {
		return "", err
	}

	return mes, nil
}

func (ser *service) Update(id int64, name string, color string, price float64, stock int64, code string, isPublished bool, createdAt string) (Product, error) {

	prod, err := ser.repository.Update(id, name, color, price, stock, code, isPublished, createdAt)

	if err != nil {
		return Product{}, err
	} else {
		return prod, nil
	}
}

func (ser *service) UpdateNombre(id int64, name string) (Product, error) {
	prod, err := ser.repository.UpdateNombre(id, name)

	if err != nil {
		return Product{}, err
	} else {
		return prod, nil
	}
}
