package store

import (
	"encoding/json"
	"os"
)

type Store interface {
	Write(data interface{}) error
	Read(data interface{}) error
}

type TypeFile string

const(
	FileType TypeFile = "file"
)

type FileStore struct {
	FileName string
}


func New(store TypeFile,filename string) Store{
	switch(store){
	case FileType:
		return &FileStore{filename}
	}

	return nil
}


func(sto *FileStore) Write(data interface{}) error{
	dataBytes, err := json.MarshalIndent(data, "", " ")

	if err != nil {
		return err
    }
	err = os.WriteFile(sto.FileName, dataBytes, 0644)
	if err != nil {
		return err
    }
	return nil
}

func(sto *FileStore) Read(data interface{}) error{

	file, err:= os.ReadFile(sto.FileName)
	if(err != nil) {
		return err
	}

	return json.Unmarshal([]byte(file), &data)
}
