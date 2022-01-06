package store

import (
	"bufio"
	"os"
)

type SaveFile interface {
	ReadLines(path string) ([]string, error)
}

type FileStoreSave struct{}

type TypeSave string

const (
	FileTypeSave TypeSave = "file"
)

func NewSave(store TypeSave) SaveFile {
	switch store {
	case FileTypeSave:
		return &FileStoreSave{}
	}
	return nil
}

func (fs *FileStoreSave) ReadLines(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}
