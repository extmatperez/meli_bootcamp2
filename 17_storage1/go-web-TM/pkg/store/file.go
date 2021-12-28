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

type FileStore struct {
	FileName string
	Mock     *Mock
}

type Mock struct {
	Data []byte
	Err  error
}

func (f *FileStore) AddMock(mock *Mock) {
	f.Mock = mock
}
func (f *FileStore) ClearMock(mock *Mock) {
	f.Mock = nil
}

func (fs *FileStore) Read(data interface{}) error {
	if fs.Mock != nil {
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
func (f *FileStore) Write(data interface{}) error { // crea un archivo y el file
	fileData, err := json.MarshalIndent(data, "", " ")
	if err != nil {
		return err
	}
	if f.Mock != nil {
		if f.Mock.Err != nil {
			return f.Mock.Err
		}
		f.Mock.Data = fileData
		return nil
	}
	return os.WriteFile(f.FileName, fileData, 0644)

}

func New(store Type, fileName string) Store {
	switch store {
	case FileType:
		return &FileStore{FileName: fileName}
	}
	return nil
}
