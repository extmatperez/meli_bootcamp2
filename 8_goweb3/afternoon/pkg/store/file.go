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
		return &FileStore{filename}
	}
	return nil
}

type FileStore struct {
	filename string
}

func (f *FileStore) Read(data interface{}) error {
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

	os.WriteFile(f.filename, fileData, 0644)

	return nil
}
