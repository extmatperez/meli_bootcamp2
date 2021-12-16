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
	FileType Type = "string"
)

type FileStore struct {
	FileName string
	Mock     *MockStore
}

type MockStore struct {
	Data []byte
	Err  error
}

func (fs *FileStore) AddMock(mock *MockStore) {
	fs.Mock = mock
}
func (fs *FileStore) ClearMock() {
	fs.Mock = nil
}

func New(store Type, fileName string) Store {
	switch store {
	case FileType:
		return &FileStore{FileName: fileName}
	}
	return nil
}

func (fs *FileStore) Write(data interface{}) error {
	//MarshalIndent me da el valor del json con formato
	dataToWrite, err := json.MarshalIndent(data, "", " ")
	if err != nil {
		return err
	}
	//Si hay mock estamos en testing
	if fs.Mock != nil {
		if fs.Mock.Err != nil {
			return fs.Mock.Err
		}
		fs.Mock.Data = dataToWrite
		return nil
	}

	//Este es el comportamiento normal sin testing
	return os.WriteFile(fs.FileName, dataToWrite, 0644)
}

func (fs *FileStore) Read(data interface{}) error {
	//Si hay mock, tomo los datos de él y no del archivo.
	if fs.Mock != nil {
		if fs.Mock.Err != nil {
			return fs.Mock.Err
		}
		return json.Unmarshal(fs.Mock.Data, &data)
	}
	//Si no hay mock hago lo mismo que antes
	file, err := os.ReadFile(fs.FileName)
	if err != nil {
		return err
	}

	return json.Unmarshal(file, &data)
}
