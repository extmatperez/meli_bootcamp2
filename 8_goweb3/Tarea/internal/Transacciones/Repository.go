package internal

type Transaccion struct {
	ID       int     `json:"id"`
	Codigo   string  `json:"codigo"`
	Moneda   string  `json:"moneda"`
	Monto    float64 `json:"monto"`
	Emisor   string  `json:"emisor"`
	Receptor string  `json:"receptor"`
}

var transaccionesTodas []Transaccion
var lastID int

type Repository interface {
	GetAll() ([]Transaccion, error)
	Store(Id int, codigo string, moneda string, monto float64, emisor string, receptor string) (Transaccion, error)
	GetID() (int, error)
}

type repository struct {
}

func NewRepository() Repository {
	return &repository{}
}

func (r *repository) GetAll() ([]Transaccion, error) {
	return transaccionesTodas, nil
}

func (r *repository) GetID() (int, error) {
	return lastID, nil
}

func (r *repository) Store(Id int, codigo string, moneda string, monto float64, emisor string, receptor string) (Transaccion, error) {
	t := Transaccion{Id, codigo, moneda, monto, emisor, receptor}
	transaccionesTodas = append(transaccionesTodas, t)
	lastID = t.ID
	return t, nil
}
