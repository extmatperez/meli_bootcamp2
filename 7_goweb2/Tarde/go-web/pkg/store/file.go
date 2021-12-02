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
}

func (fs *FileStore) Read(data interface{}) error {

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

	err = os.WriteFile(fs.FileName, content, 0644)

	if err != nil {
		return err
	}
	return nil
}

func New(store Type, fileName string) Store {
	switch store {
	case FileType:
		return &FileStore{fileName}
	}
	return nil
}
