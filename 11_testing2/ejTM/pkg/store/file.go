package store

import (
	"encoding/json"
	"os"
)

type Store interface {
	Read(data interface{}) error
	Write(data interface{}) error
}

type FileStore struct {
	FileName string
}

func NewStore(fileType string, fileName string) Store {
	switch fileType {
	case "file":
		return &FileStore{FileName: fileName}
	}
	return nil
}

func (fs *FileStore) Read(data interface{}) error {
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

	return os.WriteFile(fs.FileName, dataToWrite, 0644)
}
