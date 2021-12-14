package store

import (
	"encoding/json"
	"os"
)

type Store interface {
	Read(data interface{}) (bool, error)
	Write(data interface{}) (bool, error)
}

type Type string

const (
	FileType Type = "file_transaction"
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
}

func (fs *FileStore) AddMock(mock *Mock) {
	fs.Mock = mock
}
func (fs *FileStore) ClearMock() {
	fs.Mock = nil
}

func (fs *FileStore) Write(data interface{}) (bool, error) {
	file, err := json.MarshalIndent(data, "", " ")
	if err != nil {
		return false, err
	}
	if fs.Mock != nil {
		if fs.Mock.Err != nil {
			return false, fs.Mock.Err
		}
		fs.Mock.Data = file
		return true, nil
	}
	return true, os.WriteFile(fs.FileName, file, 0644)
}
func (fs *FileStore) Read(data interface{}) (bool, error) {
	if fs.Mock != nil {
		if fs.Mock.Err != nil {
			return false, fs.Mock.Err
		}
		return true, json.Unmarshal(fs.Mock.Data, data)
	}
	file, err := os.ReadFile(fs.FileName)
	if err != nil {
		return false, err
	}
	return true, json.Unmarshal(file, &data)
}
