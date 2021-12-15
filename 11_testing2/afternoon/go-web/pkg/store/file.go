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
