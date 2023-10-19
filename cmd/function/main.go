package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gocolly/colly"
	"github.com/thefrol/soccerway-parser-go/internal/parse"
	"github.com/thefrol/soccerway-parser-go/internal/player"
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

	var player player.Raw
	parse.Text(c, "dd[data-first_name]", &player.FirstName)
	parse.Text(c, "dd[data-last_name]", &player.LastName)
	parse.Text(c, "dd[data-nationality]", &player.Nationality)
	parse.Date(c, "dd[data-date_of_birth]", &player.DateOfBirth)
	parse.Int(c, "dd[data-age]", &player.Age)
	parse.Text(c, "dd[data-country_of_birth]", &player.CountryOfBirth)
	parse.Text(c, "dd[data-position]", &player.Position)
	parse.Text(c, "dd[data-height]", &player.Height)
	parse.Text(c, "dd[data-foot]", &player.Foot)

	err := c.Visit("http://us.soccerway.com/players/danilo-malik-santos/871010/")
	if err != nil {
		// todo.. if forbidden
		log.Fatal(err)
	}

	fmt.Printf("%+v", player)
}
