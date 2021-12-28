package internal

import "github.com/extmatperez/meli_bootcamp2/18_storage2/afternoon/mockTests/go-web/internal/models"

type ServiceSQLMock interface {
	Store(product models.Producto) (models.Producto, error)
	GetOne(id int) (models.Producto, error)
	Update(producto models.Producto, id int) (models.Producto, error)
	Delete(id int) error
}

type serviceSQLMock struct {
	repo RepositorySQLMock
}

func NewServiceSQLMock(repo RepositorySQLMock) ServiceSQLMock {
	return &serviceSQLMock{repo: repo}
}

func (ser *serviceSQLMock) Store(producto models.Producto) (models.Producto, error) {
	return ser.repo.Store(producto)
}

func (ser *serviceSQLMock) GetOne(id int) (models.Producto, error) {
	return ser.repo.GetOne(id)
}

func (ser *serviceSQLMock) Update(producto models.Producto, id int) (models.Producto, error) {
	return ser.repo.Update(producto, id)
}

func (ser *serviceSQLMock) Delete(id int) error {
	return ser.repo.Delete(id)
}
