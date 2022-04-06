package controllers

import (
	"io/ioutil"
	"log"
)

type Open interface {
	Open_txtfile() []byte
}

type Open_fields struct {
	NameText string
}

//Open text file
func (field Open_fields) Open_txtfile() []byte {
	data_io, err := ioutil.ReadFile(field.NameText)
	if err != nil {
		log.Fatal(err)
	}
	return data_io
}
