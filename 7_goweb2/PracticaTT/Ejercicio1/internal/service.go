package internal

type Service interface {
	getAll() ([]Transaccion, error)
	Store(codTransaccion, moneda string, monto float64, emisor, receptor, fechaTrans string) (Transaccion, error)
}

type Service struct {
	repository Repository
}
