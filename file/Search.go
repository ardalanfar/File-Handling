package file

import (
	"strings"
	"sync"
)

type Search interface {
	Search_word_txtfile()
}

type Search_fields struct {
	Databyt     []byte
	Word_search string
}

func (field Search_fields) Search_word_txtfile(cha chan int, wg *sync.WaitGroup) {
	result_search := strings.Count(string(field.Databyt), field.Word_search)
	cha <- result_search
	wg.Done()
}
