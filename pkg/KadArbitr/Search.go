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

	core.CloseNotification()
	// Переменная req.SearchCases хранит данные о классе, который необходимо кажать,
	// будь то обычный поиск через "Найти" или поиск в молификации гражданских, банкротств или административных делах.
	if req.SearchCases == ModeSearch {
		// Если общий запрос через кнопку "найти"
		ErrorClick = core.page.Press(`input[placeholder="фамилия судьи"]`, "Enter")
		if ErrorClick != nil {
			return ErrorClick
		}
		// ErrorClick = core.page.Click(`button[alt=Найти]`, playwright.PageClickOptions{
		// 	Force:   playwright.Bool(true),
		// 	Delay:   playwright.Float(100),
		// 	Timeout: playwright.Float(5000),
		// })
		if ErrorClick != nil {
			return ErrorClick
		}
	} else if req.SearchCases == ModeAdministrative || req.SearchCases == ModeCivil || req.SearchCases == ModeBankruptcy {
		// Если запрос по категории дела "Административные", "Гражданские", "Банкротные"
		// fmt.Println(`li[class=` + req.SearchCases + `]`)
		ErrorClick = core.page.Click(`li[class=`+req.SearchCases+`]`, playwright.PageClickOptions{
			Force:   playwright.Bool(true),
			Delay:   playwright.Float(100),
			Timeout: playwright.Float(5000),
		})
		if ErrorClick != nil {
			return ErrorClick
		}
	} else {
		// Если неправильно задан аргумент запроса
		fmt.Println(req.SearchCases)
		return ErrorAnotherSearchMode
	}

	// // Ждём ответа от POST запроса
	// core.page.WaitForResponse("https://kad.arbitr.ru/Kad/SearchInstances", playwright.FrameWaitForURLOptions{
	// 	Timeout:   playwright.Float(2),           // Таймаут на ожидание
	// 	WaitUntil: playwright.WaitUntilStateLoad, // Пока не загрузится ответ
	// 	// WaitUntil: playwright.WaitUntilStateLoad,
	// })
	core.page.WaitForSelector("div[id=table] table[class=b-cases] colgroup")

	return nil
}

// Закрыть уведомление о недоступности опубликованных аудиозаписей
func (core *CoreReq) CloseNotification() error {
	return core.page.Click(`a[class="b-promo_notification-popup-close js-promo_notification-popup-close"]`, playwright.PageClickOptions{
		Delay:   playwright.Float(100),
		Timeout: playwright.Float(3000),
	})
}
