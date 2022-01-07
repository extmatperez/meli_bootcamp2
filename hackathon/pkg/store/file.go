package store

import (
	"encoding/json"
	"os"
)

type Store interface {
	Read(data interface{}) error
	Write(data interface{}) error
}
type Mock struct {
	Data  []byte
	Error error
}
type FileStore struct {
	FileName string
	Mock     *Mock
}

func NewStore(fileType string, fileName string) Store {
	switch fileType {
	case "file":
		return &FileStore{FileName: fileName}
	}
	return nil
}
func (fs *FileStore) AddMock(mock *Mock) {
	fs.Mock = mock
}
func (fs *FileStore) ClearMock() {
	fs.Mock = nil
}

func (fs *FileStore) Read(data interface{}) error {
	if fs.Mock != nil {
		if fs.Mock.Error != nil {
			return fs.Mock.Error
		}
		return json.Unmarshal(fs.Mock.Data, &data)

	}
	file, err := os.ReadFile(fs.FileName)
	if err != nil {
		return err
	}
	return json.Unmarshal(file, &data)
}

func (fs *FileStore) Write(data interface{}) error {
	dataToWrite, err := json.MarshalIndent(data, "", " ")
	if err != nil {
		return err
	}
	if fs.Mock != nil {
		if fs.Mock.Error != nil {
			return fs.Mock.Error
		}
		fs.Mock.Data = dataToWrite
		return nil

	}

	return os.WriteFile(fs.FileName, dataToWrite, 0644)
}
