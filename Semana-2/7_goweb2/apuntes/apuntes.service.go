package internal

// el servicio deberia recibir algo de una peticion y luego pedirle al repo lo que necesite
// service va a usar los mismos metodos que estan en el repo
// no hacemos el lastID pq no nos ahce falta epro podriamos...
type Service interface {
	GetAll() ([]Persona, error) //reconoce perfectamente a []Persona pq estamos en el mismo paquete (internal)
	Store(nombre string, apellido string, email string) (Persona, error)
	// en este Store no tenemos el ID. pq en teoria cd yo creo algo no le deberia pasar el id
	//el id deberia ser un elemento  calculado aca.
	Update(id int, nombre string, apellido string, email string) (Persona, error)
}

//tambien vamos a necesitar una estructura, a la cual le vamos a pasar NO EL REPOSITORIO COMO ESTRUCTRUA,
//sino el repositorio como INTERFAZ. (PRESTAR ATENCION A LA FILE repository.go)
// de la forma en la que el servicio tiene un repo, luego el handler va a tener un service, pero el handler NO VA  a tener un repo
type service struct {
	repository Repository
}

//vamos a crear un NewService de la misma forma que creamos un NewRepository
// que va a retornar una interfaz Service
func NewService(repository Repository) Service {
	// vamos a devolver un Service (ANTES LA STRUCT REPOSITORY ESTABA VACIA PQ NO TENIA NADA. PERO AHORA.. mi sevice
	// SI tiene algo, tiene un repository, x eso lo tengo que pasar
	return &service{repository: repository}
}

//tengo que implementar los metodos

func (ser *service) GetAll() ([]Persona, error) {
	//accedo a los metodos que tiene el repositorio
	// lo meto en la variable persona, err para poder validar.
	personas, err := ser.repository.GetAll()

	if err != nil {
		return nil, err
		//devuelvo nil y error

	}
	return personas, nil
	//devuelvo personas y nil

}

// ahora implementaqmos STORE

func (ser *service) Store(nombre string, apellido string, email string) (Persona, error) {

	//aca tenemos que pasarle el id pq sino nos va a dar error arriba
	//entonces se lo pedimos al repo. IMPORTANTE!!!!
	//ACLARACION-- arriba, donde creamos la interfaz Service NO ponemos ID, pq esos son los metodos
	// que van a estar expuesrtos al handler, y el handler NO VA A NECESITAR MANIPULAR EL ID
	//pero el service si, por eso lo traemos aca
	ultimoId, err := ser.repository.LastId()
	// el ultimoID podria haber venido con error, asi que lo controlo
	if err != nil {
		return Persona{}, err
		//importantea aca, no pudee retornar un slice,  tiene que retornar una persona vacia, ya que nil no es un tipo de dato
		// apto para una estructura
	}

	//sino .....
	per, err := ser.repository.Store(ultimoId+1, nombre, apellido, email)
	// aca nunca nadie incremento el ultimo ID. por eso le metemos +1

	if err != nil {
		//validamos nuevamente . si no lopudo almacenar, devolvemos una persona vacia y un error
		return Persona{}, err
	}
	//sino, retorname la per, y nil

	return per, nil
}

func (ser *service) Update(id int, nombre, apellido, email string) (Persona, error) {
	return ser.repository.Update(id, nombre, apellido, email)
}
