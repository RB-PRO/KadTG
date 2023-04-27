package card

import "github.com/RB-PRO/KadTG/pkg/KadArbitr"

type Bases struct {
	*KadArbitr.CoreReq
}

// Файл для парсинга каждой страницы судебного дела, [например]
//
// # Используется структура Card
//
// [пример]: https://kad.arbitr.ru/Card/72197155-c243-47d3-b328-2c421391754a
func (core *Bases) ParseCard(url string) (card Carda, ErrorParse error) {

	// Переходим по ссылке с запроса
	if _, err := core.page.Goto(url); err != nil {
		return Carda{}, err // could not create page
	}

	// Статус дела
	SelectorStatus, _ := core.page.QuerySelector("div#b-case-header-desc")
	if SelectorStatus != nil { // Если найден такой блок
		// Берём текстовое значение и проверяем его на ошибку
		if FindText, IsFindError := SelectorStatus.InnerText(); IsFindError != nil {
			card.Status = FindText
		}
	}

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
