package store

import (
	"encoding/json"
	"os"
)
type TypeFile string

const(
	FileType TypeFile = "file"
)

type FileStore struct {
	FileName string
	Mock *Mock
}

type Store interface {
	Write(data interface{}) error
	Read(data interface{}) error
}

type Mock struct{
	Data []byte
	Err error
}

func (fs *FileStore) AddMock(mock *Mock){
	fs.Mock = mock
}

func (fs *FileStore) DeleteMock(mock *Mock){
	fs.Mock = nil
}
func (fs *Mock) ReadMock(Data []byte) bool{
	fs.Data = Data
	return true;
}


func New(typeFile TypeFile,filename string) Store{
	switch(typeFile){
	case FileType:
		return &FileStore{FileName:filename}
	}

	return nil
}


func(sto *FileStore) Write(data interface{}) error{
	dataBytes, err := json.MarshalIndent(data, "", " ")

	if err != nil {
		return err
    }

	if sto.Mock != nil {
		if sto.Mock.Err != nil{
			return sto.Mock.Err
		}
		 sto.Mock.Data = dataBytes
		return nil
	}


	err = os.WriteFile(sto.FileName, dataBytes, 0644)
	if err != nil {
		return err
    }
	return nil
}

func(sto *FileStore) Read(data interface{}) error{

	if sto.Mock != nil {
		if sto.Mock.Err != nil{
			return sto.Mock.Err
		}
		return json.Unmarshal(sto.Mock.Data, &data)
	}


	file, err:= os.ReadFile(sto.FileName)
	if(err != nil) {
		return err
	}

	return json.Unmarshal([]byte(file), &data)
}
