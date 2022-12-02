package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

type extractedCard struct {
	id     string
	title  string
	writer string
	count  string
}

var baseUrl string = "https://gall.dcinside.com/board/lists/?id=dcbest&_dcbest=1"

// var baseUrl string = "https://okky.kr/community?"

func main() {
	c := make(chan []extractedCard)
	var cards []extractedCard
	totalPages := getPages()

	for i := 0; i < totalPages; i++ {
		go getPage(i, c)
	}

	for i := 0; i < totalPages; i++ {
		getCards := <-c
		cards = append(cards, getCards...)
	}

	writeJobs(cards)
	fmt.Println("Done, extracted ", len(cards))
}

func writeJobs(cards []extractedCard) {
	// create writer -> input data -> save
	file, err := os.Create("cards.csv")
	checkErr(err)

	w := csv.NewWriter(file)
	defer w.Flush()

	headers := []string{"ID", "Title", "Writer", "Count"}
	wErr := w.Write(headers)
	checkErr(wErr)

	for _, card := range cards {
		cardSlice := []string{card.id, card.title, card.writer, card.count}
		cwErr := w.Write(cardSlice)
		checkErr(cwErr)
	}
}

func getPage(page int, mainC chan<- []extractedCard) {
	pageUrl := baseUrl + "&page=" + strconv.Itoa(page)
	fmt.Println(pageUrl)
	var cards []extractedCard
	res, err := http.Get(pageUrl)
	c := make(chan extractedCard)

	checkErr(err)
	checkCode(res)

	defer res.Body.Close()

	doc, err := goquery.NewDocumentFromReader(res.Body)
	checkErr(err)

	searchCards := doc.Find(".ub-content")

	searchCards.Each(func(i int, card *goquery.Selection) {
		go getCard(card, c)
	})

	for i := 0; i < searchCards.Length(); i++ {
		card := <-c
		cards = append(cards, card)
	}

	mainC <- cards
}

func getCard(card *goquery.Selection, c chan<- extractedCard) {
	id, _ := card.Attr("data-no")
	title := cleanString(card.Find(".gall_tit>a").Text())
	writer := cleanString(card.Find(".gall_writer>span").Text())
	count := cleanString(card.Find(".gall_count").Text())

	c <- extractedCard{id: id, title: title, writer: writer, count: count}
}

func getPages() int {
	pages := 0
	res, err := http.Get(baseUrl)
	checkErr(err)
	checkCode(res)

	defer res.Body.Close() // res body 는 io 라서 함수 종료 시 닫아줘야 한다. (메모리 누수 방지)

	doc, err := goquery.NewDocumentFromReader(res.Body)
	checkErr(err)

	doc.Find(".bottom_paging_box").Each(func(i int, s *goquery.Selection) {
		pages = s.Find("a").Length()
	})

	return pages
}

func checkErr(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}

func checkCode(res *http.Response) {
	if res.StatusCode != 200 {
		log.Fatalln("Request failed with Status: ", res.StatusCode)
	}
}

func cleanString(str string) string {
	// hello.   f.      1  ->  ["hello.", "f.", "1"] -> hello f 1
	return strings.Join(strings.Fields(strings.TrimSpace(str)), " ")
}

/*
								main

			getPage getPage getPage getPage getPage getPage

	getCard getCard getCard getCard getCard getCard getCard getCard getCard ...
*/
