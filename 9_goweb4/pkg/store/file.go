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
}

func (fs *FileStore) Write(data interface{}) error {
	fileData, err := json.MarshalIndent(data, "", "	")

	if err != nil {
		return err
	}

	return os.WriteFile(fs.FileName, fileData, 0644)
}

func (fs *FileStore) Read(data interface{}) error {
	fileData, err := os.ReadFile(fs.FileName)

	if err != nil {
		return err
	}

	return json.Unmarshal(fileData, &data)
}
