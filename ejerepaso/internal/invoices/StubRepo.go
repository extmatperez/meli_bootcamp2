package employee

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/extlurosell/meli_bootcamp_go_w2-2/internal/domain"
)

type StubRepository struct {
	stubRepositoryFlag bool
}

var employeeMokeService string = `[
	{"id": 1, "card_number_id": "1234561", "first_name": "Paula", "last_name": "Cabello", "warehouse_id": 3},
	{"id": 2, "card_number_id": "1234562", "first_name": "Jose", "last_name": "Rios", "warehouse_id": 3}]`

/*
FUNCIONES MOKEADAS DE REPOSITORY
*/

/*
Funcion para obtener todos los empleados, estos empleados se agregan a una
lista llamada employeeMokeService, los cuales se transforman a una lista de
Employees. la cual se devuelve y se muestra.
*/
func (s *StubRepository) GetAll(ctx context.Context) ([]domain.Employee, error) {
	var listEmployeeOut []domain.Employee
	err := json.Unmarshal([]byte(employeeMokeService), &listEmployeeOut)
	s.stubRepositoryFlag = true
	return listEmployeeOut, err
}

/*
Con esta funcion se consulta por un ID dentro de la lista de employeeMokeService
en el caso de que un ID coincida, se retorna este valor con un error nil
y en caso contrario se retorna un employee vacio y un error de employee not found
*/
func (s *StubRepository) Get(ctx context.Context, id int) (domain.Employee, error) {
	var employee domain.Employee
	listEmployeeOut, err := s.GetAll(context.Background())
	for _, emp := range listEmployeeOut {
		if emp.ID == id {
			employee = emp
			err = nil
			s.stubRepositoryFlag = true
			return employee, err
		}
	}
	s.stubRepositoryFlag = true
	return domain.Employee{}, fmt.Errorf("Employee not found")

}

/*
Funcion que retorna un Employee con los datos de ingreso y retorna el ID ingresad.
*/
func (s *StubRepository) Save(ctx context.Context, empl domain.Employee) (int, error) {
	newEmployee := domain.Employee{}
	err := fmt.Errorf("")
	if empl.CardNumberID == "" || empl.FirstName == "" || empl.LastName == "" || empl.WarehouseID == 0 {
		err = fmt.Errorf("Atribute not found")
	} else {
		emplExists := s.Exists(context.Background(), empl.CardNumberID)
		if !emplExists {
			newEmployee.ID = empl.ID
			newEmployee.CardNumberID = empl.CardNumberID
			newEmployee.FirstName = empl.FirstName
			newEmployee.LastName = empl.LastName
			newEmployee.WarehouseID = empl.WarehouseID

			err = nil
		}
	}
	s.stubRepositoryFlag = true
	return newEmployee.ID, err
}

/*
Funcion update la cual recibe un employee y consulta si el id de este employee existe,
en el caso de que exista retorna un error = nil y en el caso de que no exista
retorna un Error = Employee not found
*/
func (s *StubRepository) Update(ctx context.Context, empl domain.Employee) error {
	_, err := s.Get(context.Background(), empl.ID)
	s.stubRepositoryFlag = true
	return err
}

/*
Funcion delete la cual recibe un id al cual pregunta si existe.
en el caso de que exista retorna un error = nil y en el caso de que no exista
retorna un Error = Employee not found
*/
func (s *StubRepository) Delete(ctx context.Context, id int) error {
	_, err := s.Get(context.Background(), id)
	s.stubRepositoryFlag = true
	return err
}

/*
Funcion que retorna un true en caso de que el cardNumberID exista en
la lista de employeeMokeService, en caso contrario retorna un false.
*/
func (s *StubRepository) Exists(ctx context.Context, cardNumberID string) bool {
	var exist bool = false
	listEmployeeOut, _ := s.GetAll(context.Background())
	for _, emp := range listEmployeeOut {
		if emp.CardNumberID == cardNumberID {
			exist = true
		}
	}
	s.stubRepositoryFlag = true
	return exist
}
