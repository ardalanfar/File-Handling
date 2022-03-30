package main

import (
	"Edit_File/conf"
	"Edit_File/file"
	"Edit_File/scrap"
	"flag"
	"fmt"
	"sync"
)

func main() {

	config := conf.GetConfig()

	WordSearch := flag.String("ws", "nil", "word to serch in file")
	NumberLines := flag.Int("nl", 0, "number of lines")
	Replace_oldWord := flag.String("rplo", "nil", "old word")
	Replace_newWord := flag.String("rpln", "nil", "new word")
	flag.Parse()

	if *NumberLines != 0 || *WordSearch != "nil" || (*Replace_newWord != "nil" && *Replace_oldWord != "nil") {

		/*----------------------------------------------------*/
		//Scrap Target

		wg_scraplink := sync.WaitGroup{}

		scarp_t := scrap.Scrap_fields{
			Web_address:    config.Web_address,
			Allowed_domain: config.Allowed_domain,
		}
		strlink := scarp_t.Scrap_links()

		save_scrap := file.Save_fields{
			Databyt:  []byte(strlink),
			NameText: config.Address_text_file,
		}

		wg_scraplink.Add(1)
		cha_scraplink := make(chan bool)
		go save_scrap.Save_txtfile(cha_scraplink, &wg_scraplink)
		<-cha_scraplink
		wg_scraplink.Wait()

		/*----------------------------------------------------*/
		//Open file

		open_f := file.Open_fields{
			NameText: config.Address_text_file,
		}

		databyt := open_f.Open_txtfile()

		wg_search := sync.WaitGroup{}
		wg_numberl := sync.WaitGroup{}
		wg_save := sync.WaitGroup{}

		/*----------------------------------------------------*/
		//Search word

		if *WordSearch != "nil" {

			search_f := file.Search_fields{
				Word_search: *WordSearch,
				Databyt:     databyt,
			}

			wg_search.Add(1)
			cha_search := make(chan int)
			go search_f.Search_word_txtfile(cha_search, &wg_search)
			result_s := <-cha_search

			fmt.Println("Number Of Searched Word :", result_s)
		}

		/*----------------------------------------------------*/
		// Counting lines

		if *NumberLines != 0 {

			counting_f := file.Counting_fields{
				Databyt: databyt,
			}

			wg_numberl.Add(1)
			cha_numberlines := make(chan int)
			go counting_f.Number_lines_textfile(cha_numberlines, &wg_numberl)
			result_nl := <-cha_numberlines

			fmt.Println("Number Lines File : ", result_nl)
		}

		/*----------------------------------------------------*/
		//Replace word

		if *Replace_newWord != "nil" && *Replace_oldWord != "nil" {

			replace_f := file.Replace_fields{
				Databyt:       databyt,
				Rplc_old_word: *Replace_oldWord,
				Rplc_new_word: *Replace_newWord,
			}

			save_f := file.Save_fields{
				Databyt:  replace_f.Replace_word_txtfile(),
				NameText: config.Address_text_file,
			}

			wg_save.Add(1)
			cha_replace := make(chan bool)
			go save_f.Save_txtfile(cha_replace, &wg_save)
			result_rpl := <-cha_replace

			if result_rpl {
				fmt.Println("Replace ok")
			}
		}

		/*----------------------------------------------------*/

		wg_numberl.Wait()
		wg_search.Wait()
		wg_save.Wait()
	}
}
