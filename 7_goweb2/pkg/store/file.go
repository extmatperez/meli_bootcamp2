/*
	Author: AG-Meli - Andrés Ghione
*/

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

// Read
// @Summary Read from specified file
// @Tags Transaction
// @Description Read from specified file
func (f FileStore) Read(data interface{}) error {
	file, err := os.ReadFile(f.FileName)
	if err != nil {
		return err
	}
	return json.Unmarshal(file, &data)
}

// Write
// @Summary Overwrite the indicated file
// @Tags Transaction
// @Description Overwrite the indicated file
func (f FileStore) Write(data interface{}) error {
	fileData, err := json.MarshalIndent(data, "", "   ")
	if err != nil {
		return err
	}
	return os.WriteFile(f.FileName, fileData, 0644)
}

func New(store Type, fileName string) Store {
	switch store {
	case FileType:
		return &FileStore{fileName}
	}
	return nil
}
