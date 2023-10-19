package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gocolly/colly"
)

// Hanler это точка входа для веб-контейнера типа aws-lambda,
// получает в url ссылку, и выдает в теле ответа джейсон с данными
func Handler(w http.ResponseWriter, r *http.Request) {

	w.Write([]byte("hello, parser"))
}

// Эта функция не запускает сейчас локальный сервер, а просто исползуется
// как сендбокс
func main() {
	c := colly.NewCollector(colly.UserAgent("Mozilla/5.0 (X11; Linux x86_64; rv:109.0) Gecko/20100101 Firefox/117.0"))

	// Find and visit all links
	c.OnHTML("dd[data-first_name]", func(e *colly.HTMLElement) {
		fmt.Println(e.Text)
	})

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL)
	})

	err := c.Visit("http://us.soccerway.com/players/danilo-malik-santos/871010/")
	if err != nil {
		// todo.. if forbidden
		log.Fatal(err)
	}
}
