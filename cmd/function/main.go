package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gocolly/colly"
	"github.com/thefrol/soccerway-parser-go/internal/parse"
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

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL)
	})

	var firstName, lastName string
	parse.Text(c, "dd[data-first_name]", &firstName)
	parse.Text(c, "dd[data-last_name]", &lastName)

	err := c.Visit("http://us.soccerway.com/players/danilo-malik-santos/871010/")
	if err != nil {
		// todo.. if forbidden
		log.Fatal(err)
	}

	fmt.Println(firstName, lastName)
}
