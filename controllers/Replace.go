package controllers

import (
	"bytes"
)

type Replace interface {
	Replace_word_txtfile() []byte
}

type Replace_fields struct {
	Databyt       []byte
	Rplc_new_word string
	Rplc_old_word string
}

//Replace word in text file
func (field Replace_fields) Replace_word_txtfile() []byte {
	rep_newword := " " + field.Rplc_new_word + " "
	rep_oldword := " " + field.Rplc_old_word + " "

	result_rpl := bytes.ReplaceAll(field.Databyt, []byte(rep_oldword), []byte(rep_newword))
	return result_rpl
}
