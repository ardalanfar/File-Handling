package file

import (
	"strings"
	"sync"
)

type Counting interface {
	Number_lines_textfile()
}

type Counting_fields struct {
	Databyt []byte
}

func (field Counting_fields) Number_lines_textfile(cha chan int, wg *sync.WaitGroup) {
	var sum int
	lines := strings.Split(string(field.Databyt), "\n")
	for sum = range lines {
		sum = sum + 1
	}
	cha <- sum
	wg.Done()
}
