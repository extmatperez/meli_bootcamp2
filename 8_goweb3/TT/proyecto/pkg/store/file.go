/*
Se debe implementar la funcionalidad para guardar la información de la petición en un
archivo json, para eso se deben realizar los siguientes pasos:
	1. En lugar de guardar los valores de nuestra entidad en memoria, se debe crear un
	archivo; los valores que se vayan agregando se guardan en él.
*/

/*
Se debe implementar la funcionalidad para leer la información requerida en la petición del
archivo json generado al momento de guardar, para eso se deben realizar los siguientes
pasos:
	1. En lugar de leer los valores de nuestra entidad en memoria, se debe obtener del 
	archivo generado en el punto anterior.
*/

package store

import (
	"encoding/json"
	"os"
)

type Store interface {
	Read(data interface{}) error
	Write(data interface{}) error
}

type Type string

const (
	FileType Type = "file"
)

func New(store Type, fileName string) Store {

	switch store {
	case FileType:
		return &FileStore{fileName}
	}

	return nil
}

type FileStore struct {
	FileName string
}

func (fs *FileStore) Write(data interface{}) error {

	file, err := json.MarshalIndent(data, "", " ")

	if err != nil {
		return err
	}

	return os.WriteFile(fs.FileName, file, 0644)
}

func (fs *FileStore) Read(data interface{}) error{

	file, err := os.ReadFile(fs.FileName)

	if err != nil {
		return err
	}

	return json.Unmarshal(file, &data)
}