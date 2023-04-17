package KadArbitr

import (
	"errors"
	"fmt"

	"github.com/playwright-community/playwright-go"
)

// Ошибка при которой неверно указана модификация поиска
// Исправьте поле SearchCases в структуре Request, чтобы
var ErrorAnotherSearchMode error = errors.New("Search: Вы выбрали неверное название модификатора поиска")

// ## Нажать на кнопку поиска
//
// KadArbitr работает интересным образом. Поиск можно запустить, нажав на кнопки из верхнего меню а,г,б.
// Однако если мы хотим получить инфомацию по всем таким делам, то нам нужно нажать на кнопку "найти",
// которая иногда работает не совсем корректно, открывая поле выбора даты "По"
// Поэтому мы будем эмулировать нажатие кнопки Enter по полю номера дела, чтобы запустить поиск для общих случаев.
//
// Выбрать, какой будет запрос. Варианты:
//   - administrative active - Административные
//   - civil - Гражданские
//   - bankruptcy - Банкротсные
//   - "" - Пустой
//
// Эти занчения заданы с помощью констант:
//   - ModeAdministrative string = "administrative" // Административные - Параметры поиска
//   - ModeCivil          string = "civil"          // Гражданские - Параметры поиска
//   - ModeBankruptcy     string = "bankruptcy"     // Банкротные - Параметры поиска
//   - ModeSearch         string = ""               // Найти. Обычный поиск. - Параметры поиска
func (core *CoreReq) Search(req Request) (ErrorClick error) {
	// Переменная req.SearchCases хранит данные о классе, который необходимо кажать,
	// будь то обычный поиск через "Найти" или поиск в молификации гражданских, банкротств или административных делах.
	if req.SearchCases == ModeSearch {
		// Если общий запрос через кнопку "найти"
		ErrorClick = core.page.Press(`div[id=sug-cases] div[class=tag] input[placeholder="например, А50-5568/08"]`, "Enter")
		if ErrorClick != nil {
			return ErrorClick
		}
	} else if req.SearchCases == ModeAdministrative || req.SearchCases == ModeCivil || req.SearchCases == ModeBankruptcy {
		// Если запрос по категории дела "Административные", "Гражданские", "Банкротные"
		core.page.Click(`li[class=` + req.SearchCases + `]`)
		if ErrorClick != nil {
			return ErrorClick
		}
	} else {
		// Если неправильно задан аргумент запроса
		fmt.Println(req.SearchCases)
		return ErrorAnotherSearchMode
	}

	// Ждём ответа от POST запроса
	core.page.WaitForResponse("https://kad.arbitr.ru/Kad/SearchInstances", playwright.FrameWaitForURLOptions{
		Timeout:   playwright.Float(2),           // Таймаут на ожидание
		WaitUntil: playwright.WaitUntilStateLoad, // Пока не загрузится ответ
	})

	// // save
	// html, _ := core.page.QuerySelector("body")
	// htmlB, _ := html.InnerHTML()
	// ioutil.WriteFile("output.html", []byte(htmlB), 0644)
	return nil
}
