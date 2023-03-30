// Этот код получает список всех судов и его значение для запроса

package KadArbitr

import (
	"fmt"

	"github.com/gocolly/colly/v2"
)

func Couters() (map[string]string, error) {
	c := colly.NewCollector()
	var ErrorURL error
	couters := make(map[string]string)

	// Поиск данных
	c.OnHTML("select[class='select js-select'] option", func(e *colly.HTMLElement) {

		value, valueExit := e.DOM.Attr("value")
		if valueExit { // Если существует такой аттрибут
			if e.DOM.Text() != "" { // Если текст не нулевой
				couters[e.DOM.Text()] = value
			}
		}

		fmt.Println(e.DOM.Text(), value)

	})

	c.OnHTML("body", func(e *colly.HTMLElement) {

		fmt.Println(e.DOM.Text())

	})

	// Делаем запрос
	ErrorURL = c.Visit(URL)
	if ErrorURL != nil {
		return nil, ErrorURL
	}

	return couters, nil
}
