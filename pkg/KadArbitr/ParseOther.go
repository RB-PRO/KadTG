package KadArbitr

import (
	"strings"

	"github.com/playwright-community/playwright-go"
)

// Файл для парсинга каждой страницы судебного дела, [например]
//
// # Используется структура Card
//
// [пример]: https://kad.arbitr.ru/Card/72197155-c243-47d3-b328-2c421391754a
func (core *CoreReq) ParseCard(url string) (card Card, ErrorParse error) {

	// Переходим по ссылке с запроса
	if _, err := core.page.Goto(url); err != nil {
		return Card{}, err // could not create page
	}

	core.Screen("screens/Card2.jpg")
	// Ждём загрузку определённой части страницы
	_, ErrorWait := core.page.WaitForSelector("dd[id=main-column]", playwright.PageWaitForSelectorOptions{Timeout: playwright.Float(20000)})
	if ErrorWait != nil {
		return Card{}, ErrorWait
	}

	// Статус дела
	SelectorStatus, _ := core.page.QuerySelector(`dt[class^=b-iblock__header]`)
	if SelectorStatus != nil { // Если найден такой блок
		// Берём текстовое значение и проверяем его на ошибку
		if FindText, IsFindError := SelectorStatus.TextContent(); IsFindError == nil {
			card.Type = strings.TrimSpace(FindText)
		}
	}

	// Следующее заседание
	// b-instanceAdditional

	core.Screen("screens/Card3.jpg")

	// // Сперва пропарсим главные значения карточек
	// MainsH, err := core.page.QuerySelectorAll(`div[class="b-chrono-item-header js-chrono-item-header page-break"]`)
	// if err != nil {
	// 	return card, err // could not get entries
	// }
	// // Выделяем память для карточек
	// card.Slips = make([]struct {
	// 	Main  HistoryMain
	// 	Slave []HistorySlave
	// }, len(MainsH))
	// for IndexMain, mainH := range MainsH {
	// 	card.Slips[IndexMain].Main.
	// }

	return card, ErrorParse
}
