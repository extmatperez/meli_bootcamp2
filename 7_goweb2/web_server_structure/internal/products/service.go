package internal

type Service struct {
	repository Repository
}

type Service interface {
	getAll()
	getProductbyID()
	addProduct()
}
