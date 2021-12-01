package internal

type Service interface{
	GetAll() ([]Transaction, error)
	Store(codigo, moneda , monto, emisor, receptor,fecha string) (Transaction, error)
}

type service struct{
	repository Repository
}

func NewService(repo Repository) Service{
	return &service{repo}
}


func (ser *service) GetAll() ([]Transaction, error) {

	transactions,err := ser.repository.GetAll()

	if(err != nil){
		return nil,err
	}

	return transactions,nil
}

func (ser *service) Store(codigo, moneda , monto, emisor, receptor,fecha string) (Transaction, error) {
	
	lastID,err := ser.repository.LastId()

	if(err != nil){
		return Transaction{},err
	}

	lastID++
	transcation, err1 := ser.repository.Store(lastID,codigo,moneda,monto,emisor,receptor,fecha)
	
	if(err1 != nil){
		return Transaction{},err
	}



	return transcation,err1
}