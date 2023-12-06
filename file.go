package main

import (
	"os"
)

type file struct {
}

type File interface {
	Create(fileName string, data []byte, perm uint32) error
	Read(fileName string) (data []byte, err error)
	Remove(fileName string) error
}

func (f file) Create(fileName string, data []byte, perm uint32) error {
	err := os.WriteFile(fileName, data, os.FileMode(perm))

	if err != nil {
		return err
	}
	return nil
}

func (f file) Read(fileName string) ([]byte, error) {
	data, err := os.ReadFile(fileName)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (f file) Remove(fileName string) error {
	err := os.Remove(fileName)
	if err != nil {
		return err
	}
	return nil
}

func New() File {
	return &file{}
}
