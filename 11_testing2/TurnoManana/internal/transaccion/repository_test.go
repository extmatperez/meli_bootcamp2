package internal

import (
	"encoding/json"
	"testing"
	"github.com/stretchr/testify/assert"
)


var Datos string =  `[{
	"id": 2,
	"codigo": "24safdsadfasdf385",
	"moneda": "Peso Colombiano",
	"monto": "$8228845654645678",
	"emisor": "Luis",
	"receptor": "Perez",
	"fecha": "01/01/2001"
   },
   {
	"id": 3,
	"codigo": "11673-417",
	"moneda": "Franc",
	"monto": "$2.76",
	"emisor": "minstone2",
	"receptor": "sinnott2",
	"fecha": "1/4/2021"
   }]`


type StubStore struct{
	useMethodRed bool
}

func(s *StubStore) Write(data interface{}) error{
	return nil
}
func(s *StubStore) Read(data interface{}) error{
	s.useMethodRed = true
	return json.Unmarshal([]byte(Datos), &data)
}

func TestGetAll(t *testing.T){
	stubStore := &StubStore{}
	repo := NewRepository(stubStore)



	var excepted []Transaction
	json.Unmarshal([]byte(Datos), &excepted)

	tran,err := repo.GetAll()

	assert.Nil(t,err)
	assert.Equal(t,excepted,tran)
}


func TestUpdateName(t *testing.T){
	stubStore := &StubStore{false}
	repo := NewRepository(stubStore)
	tran2,_ := repo.GetTransactionById(2)
	codgUpdate :="AfterUpdatecod-123"


	tranUpdate,err := repo.Update(2,codgUpdate,"Peso","55.6","pepe","luis","13/12/2021")

	assert.True(t,stubStore.useMethodRed)
	assert.Equal(t,tran2.ID,tranUpdate.ID)
	assert.Equal(t,codgUpdate, codgUpdate)
	assert.Nil(t,err)
}