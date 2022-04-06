package controllers

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

//Count number of lines text file
func (field Counting_fields) Number_lines_textfile(cha chan int, wg *sync.WaitGroup) {
	var sumline int
	lines := strings.Split(string(field.Databyt), "\n")
	for sumline = range lines {
		sumline = sumline + 1
	}
	cha <- sumline
	wg.Done()
}
