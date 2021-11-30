package internal

type Transaccion struct {
	Id             int     `json:"id"`
	CodTransaccion string  `json:"cod_transaccion"`
	Moneda         string  `json:"moneda"`
	Monto          float64 `json:"monto"`
	Emisor         string  `json:"emisor"`
	Receptor       string  `json:"receptor"`
	FechaTrans     string  `json:"fecha_trans"`
}

var transacciones []Transaccion
var lastID int

type Repository interface {
	getAll() ([]Transaccion, error)
	Store(id int, codTransaccion, moneda string, monto float64, emisor, receptor, fechaTrans string) (Transaccion, error)
	LastId() (int, error)
}
