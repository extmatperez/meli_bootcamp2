package internal

type Service interface{
	GetAll() ([]Transaccion, error)
	Store(codigo, moneda , monto, emisor, receptor,fecha string) (Transaccion, error)
}

type service struct{
	repository Repository
}

func NewService(repo Repository) Service{
	return &service{repo}
}


func (ser *service) GetAll() ([]Transaccion, error) {

	transactions,err := ser.repository.GetAll()

	if(err != nil){
		return nil,err
	}

	return transactions,nil
}

func (ser *service) Store(codigo, moneda , monto, emisor, receptor,fecha string) (Transaccion, error) {
	
	lastID,err := ser.repository.LastId()

	if(err != nil){
		return Transaccion{},err
	}

	lastID++
	transcation, err1 := ser.repository.Store(lastID,codigo,moneda,monto,emisor,receptor,fecha)
	
	if(err1 != nil){
		return Transaccion{},err
	}



	return transcation,err1
}