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

type FileStore struct {
	FileName string
	Mock     *MockStore
}

type MockStore struct {
	Data    []byte
	Invoked bool
	Err     error
}

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

func (fs *FileStore) Read(data interface{}) error {

	if fs.Mock != nil {
		if fs.Mock.Err != nil {
			return fs.Mock.Err
		}
		fs.Mock.Invoked = true
		return json.Unmarshal(fs.Mock.Data, &data)
	}

	file, err := os.ReadFile(fs.FileName)
	if err != nil {
		return err
	}
	return json.Unmarshal(file, &data)
}

func (fs *FileStore) Write(data interface{}) error {

	if fs.Mock != nil {
		if fs.Mock.Err != nil {
			return fs.Mock.Err
		}
		//aqui traigo la data actualizada y la deberia grabar en la bd mockeada, lo guardo en la data solo para hacer algo...
		//no puedo guardarlo en la bd mockeada porque se hace una llamada ciclica, service_test llama a este package y por lo tanto, este no puede llamar al package internal(donde esta service_test), se podria almacenar la bd mockeada en un tercer package o en un json y cambiar la logica para que ambos packages puedan acceder a la misma y hacer persistir los datos, en fin:
		dataBytes, err := json.Marshal(data)

		fs.Mock.Data = dataBytes
		return err
	}

	fileData, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(fs.FileName, fileData, 0644)
}
