package store

import (
	"encoding/json"
	"os"
)

type Store interface {
	Write(data interface{}) error
	Read(data interface{}) error
}

type Type string

const (
	FileType Type = "file"
)

func New(store Type, fileName string) Store {
	switch store {
	case FileType:
		return &FileStore{
			FileName: fileName,
		}
	}

	return nil
}

type FileStore struct {
	FileName string
	Mock     *Mock
}

func (fs *FileStore) Write(data interface{}) error {
	fileData, err := json.MarshalIndent(data, "", "	")

	if err != nil {
		return err
	}

	if fs.Mock != nil {
		if fs.Mock.Data != nil {
			return fs.Mock.Error
		}

		fs.Mock.Data = fileData
		return nil
	}

	return os.WriteFile(fs.FileName, fileData, 0644)
}

func (fs *FileStore) Read(data interface{}) error {
	if fs.Mock != nil {
		fs.Mock.ReadUsed = true
		if fs.Mock.Error != nil {

			return fs.Mock.Error
		}

		return json.Unmarshal(fs.Mock.Data, &data)
	}

	fileData, err := os.ReadFile(fs.FileName)

	if err != nil {
		return err
	}

	return json.Unmarshal(fileData, &data)
}

func (fs *FileStore) AddMock(mock *Mock) {
	fs.Mock = mock
}

func (fs *FileStore) ClearMock() {
	fs.Mock = nil
}

type Mock struct {
	Data     []byte
	Error    error
	ReadUsed bool
}
