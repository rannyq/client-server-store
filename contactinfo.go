package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type FileStorage struct {
}

type ContactInfo struct {
	ID     string
	Name   string
	Street string
	City   string
	Zip    string
	Phone  string
}

type Storage interface {
	Add(entityID string, contactinfo ContactInfo) error
}

func (f *FileStorage) Add(entityID string, contactinfo ContactInfo) error {

	filename := fmt.Sprintf("%s.json", entityID)

	jsonBytes, err := json.Marshal(contactinfo)

	if err != nil {
		return err
	}

	err = ioutil.WriteFile(filename, jsonBytes, 0644)

	if err != nil {
		return err
	}
	return nil
}

func main() {

}
