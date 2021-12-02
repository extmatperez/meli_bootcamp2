package store

type Store interface {
	Write(data interface{})
	Read(data interface{})
}

type TypeFile string

const(
	FileName TypeFile = ""
)



func New(ruta TypeFile,) Store{

}


