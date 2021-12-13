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

func NewStore(store Type, filename string) Store {
	switch store {
	case FileType:
		return &FileStore{filename: filename}
	}
	return nil
}

type FileStore struct {
	filename string
	Mock     *Mock
}

type Mock struct {
	Data      []byte
	Err       error
	EnterRead bool
}

func (f *FileStore) AddMock(mock *Mock) {
	f.Mock = mock
}

func (f *FileStore) ClearMock() {
	f.Mock = nil
}

func (f *FileStore) Read(data interface{}) error {
	if f.Mock != nil {
		f.Mock.EnterRead = true
		if f.Mock.Err != nil {
			return f.Mock.Err
		}

		return json.Unmarshal(f.Mock.Data, &data)
	}

	fileData, err := os.ReadFile(f.filename)
	if err != nil {
		return err
	}
	json.Unmarshal(fileData, &data)
	return nil
}

func (f *FileStore) Write(data interface{}) error {
	fileData, err := json.MarshalIndent(data, "", "\t")
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

	os.WriteFile(f.filename, fileData, 0644)

	return nil
}
