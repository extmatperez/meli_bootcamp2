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

func (fs *FileStore) AddMock(mock *Mock) {
	fs.Mock = mock
}

func (fs *FileStore) ClearMock() {
	fs.Mock = nil
}

func (fs *FileStore) Read(data interface{}) error {

	if fs.Mock != nil {
		if fs.Mock.Err != nil {
			return fs.Mock.Err
		}
		return json.Unmarshal(fs.Mock.Data, data)
	}

	content, err := os.ReadFile(fs.FileName)

	if err != nil {
		return err
	}
	err = json.Unmarshal(content, data)
	if err != nil {
		return err
	}
	return nil
}

func (fs *FileStore) Write(data interface{}) error {

	content, err := json.Marshal(data)

	if err != nil {
		return err
	}

	if fs.Mock != nil {
		if fs.Mock.Err != nil {
			return fs.Mock.Err
		}
		fs.Mock.Data = content
		return nil
	}

	err = os.WriteFile(fs.FileName, content, 0644)

	if err != nil {
		return err
	}
	return nil
}

func New(store Type, fileName string) Store {
	switch store {
	case FileType:
		return &FileStore{FileName: fileName}
	}
	return nil
}
