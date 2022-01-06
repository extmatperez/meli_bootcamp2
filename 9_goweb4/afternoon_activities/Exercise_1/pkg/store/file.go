package store

import (
	"encoding/json"
	"os"
)

// Creo la interface que va a recibir los métodos con los cuales voy a leer y escribir el archivo seleccionado (CSV, JSON, etc)
type Store interface {
	Read(data interface{}) error
	Write(data interface{}) error
}

// Implementamos una constante de tipo Type para definir el tipo de store que se utilizará (File_type)
type Type string

const (
	File_type Type = "file"
)

// Implementamos función Factory(New) que se va a encargar de generar la estructura que deseamos y recibe el tipo store que queremos implementar
// y el nombre del archivo
func New(store Type, file_name string) Store {
	switch store {
	case File_type:
		return &File_store{file_name}
	}
	return nil
}

// Se declara la estructura File_store con el campo que guarde el nombre del archivo
type File_store struct {
	File_name string
}

// Declaramos el método para escribir datos en la estructura JSON. Recibe una interface y la convierte a una representación JSON en bytes para guardar
// en el archivo que especificamos al momento de instanciar la función Factory(New)
func (fs *File_store) Write(data interface{}) error {
	file, err := json.MarshalIndent(data, "", " ")
	if err != nil {
		return err
	}
	return os.WriteFile(fs.File_name, file, 0644)
}

// Declaramos el método para leer datos implementando la interface que recibe como parámetro
func (fs *File_store) Read(data interface{}) error {
	file, err := os.ReadFile(fs.File_name)

	if err != nil {
		return err
	}
	return json.Unmarshal(file, &data)
}
