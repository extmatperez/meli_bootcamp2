package store

import (
	"encoding/json"
	"os"
)

type Store interface {
	Read(data interface{}) error
	Write(data interface{}) error
}

//factory
type Type string

const (
	FileType Type = "file"
)

func NewStore(store Type, fileName string) Store {
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
	Err  error
	Data []byte
}

func (fileStore *FileStore) Read(data interface{}) error {

	// _, err := os.ReadFile("/dbProductos.json")
	// if err != nil {
	// marshaleado , err := json.Marshal([])
	// os.WriteFile(fileStore.FileName, file, 0644)
	// }
	if fileStore.Mock != nil {
		if fileStore.Mock.Err != nil {
			return fileStore.Mock.Err
		}
		return json.Unmarshal(fileStore.Mock.Data, &data)
	}
	file, err := os.ReadFile(fileStore.FileName)
	if err != nil {
		return err
	}
	return json.Unmarshal(file, &data)
}

func (fileStore *FileStore) Write(data interface{}) error {
	file, err := json.MarshalIndent(data, "", " ")
	if err != nil {
		return err
	}
	return os.WriteFile(fileStore.FileName, file, 0644)
}
