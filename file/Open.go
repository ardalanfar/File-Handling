package file

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

func (field Open_fields) Open_txtfile() []byte {
	data_io, err_io := ioutil.ReadFile(field.NameText)
	if err_io != nil {
		log.Fatal(err_io)
	}
	return data_io
}
