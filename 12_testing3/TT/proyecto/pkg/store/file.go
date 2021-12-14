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
		return &FileStore{FileName: fileName}
	}

	return nil
}

type FileStore struct {
	FileName string
	Mock     *Mock
}

type Mock struct {
	Data []byte
	Err  error
	Used bool
}

func (fs *FileStore) AddMock(mock *Mock) {
	fs.Mock = mock
}

func (fs *FileStore) ClearMock() {
	fs.Mock = nil
}

func (fs *FileStore) Write(data interface{}) error {

	file, err := json.MarshalIndent(data, "", " ")

	if err != nil {
		return err
	}

	if fs.Mock != nil {

		fs.Mock.Used = true

		if fs.Mock.Err != nil {
			return fs.Mock.Err
		}

		fs.Mock.Data = file
		return nil
	}

	return os.WriteFile(fs.FileName, file, 0644)
}

func (fs *FileStore) Read(data interface{}) error {

	if fs.Mock != nil {

		fs.Mock.Used = true

		if fs.Mock.Err != nil {
			return fs.Mock.Err
		}

		return json.Unmarshal(fs.Mock.Data, data)
	}

	file, err := os.ReadFile(fs.FileName)

	if err != nil {
		return err
	}

	return json.Unmarshal(file, &data)
}
