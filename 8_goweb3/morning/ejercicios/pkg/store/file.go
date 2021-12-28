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
	Mock *Mock
}

type Mock struct {
	Data []byte
	Err error
}

func New(store Type, fileName string) Store {
	switch store {
	case FileType:
		return &FileStore{FileName: fileName}
	}
	return nil
}

func (fs *FileStore) Write(data interface{}) error {
	file, err := json.MarshalIndent(data, "", " ")
	if err != nil {
		return err
	}
	return os.WriteFile(fs.FileName, file, 0644)
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