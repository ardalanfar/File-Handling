package file

import (
	"io/ioutil"
	"log"
	"os"
	"sync"
)

type Save interface {
	Save_txtfile()
}

type Save_fields struct {
	Databyt  []byte
	NameText string
}

func (field Save_fields) Save_txtfile(cha chan bool, wg *sync.WaitGroup) {
	file_stat, _ := os.Stat(field.NameText)

	if file_stat == nil {
		_, err_create := os.Create(field.NameText)
		if err_create != nil {
			log.Fatal(err_create)
		}
	}
	err := ioutil.WriteFile(field.NameText, field.Databyt, 0777)
	if err != nil {
		log.Fatal(err)
	}

	cha <- true
	wg.Done()
}
