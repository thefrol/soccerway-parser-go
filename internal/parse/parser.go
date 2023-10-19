// В этом пакете лежит основная логика парсинга,
// классны и функции, которые позволяют доставать
// из HTML данные и преобразовывать их
package parse

import (
	"log"
	"strconv"
	"time"

	"github.com/gocolly/colly"
)

// Text достает текстовое значение тега, указанного в tag, и записывает
// в переменную, на которую ссылкается out. Данные беруться из коллектора c
// Появятся в переменных, когда отработает коллектор.
//
// По сути эта переменная добавляет коллбек в коллектор с, этот коллбек отработает
// когда будет запущена c.Visit(url)
func Text(c *colly.Collector, tag string, out *string) {
	c.OnHTML(tag, func(e *colly.HTMLElement) {
		*out = e.Text
	})
}

// Int достает переменную типа int из тега tag. Работает
// как Text(), только преобразует строку в число
func Int(c *colly.Collector, tag string, out *int) {
	c.OnHTML(tag, func(e *colly.HTMLElement) {
		i, err := strconv.ParseInt(e.Text, 10, 64)
		if err != nil {
			log.Printf("cant parse %v as int", tag)
		}
		*out = int(i)
	})
}

// Date достает переменную типа time.Time из тега tag. Работает
// как Text(), только преобразует строку в число
func Date(c *colly.Collector, tag string, out *time.Time) {
	c.OnHTML(tag, func(e *colly.HTMLElement) {
		t, err := time.Parse("2 January 2006", e.Text)
		if err != nil {
			log.Printf("cant parse %v as time", tag)
		}
		*out = t
	})
}
