package KadArbitr

import (
	"fmt"
	"strings"
	"time"

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

	// core.Screen("screens/Card2.jpg")
	// Ждём загрузку определённой части страницы
	_, ErrorWait := core.page.WaitForSelector("dd[id=main-column]", playwright.PageWaitForSelectorOptions{Timeout: playwright.Float(20000)})
	if ErrorWait != nil {
		return Card{}, ErrorWait
	}

	// Тип дела
	if Selector, _ := core.page.QuerySelector(`dt[class^=b-iblock__header] span`); Selector != nil { // Если найден такой блок
		// Берём текстовое значение и проверяем его на ошибку
		if FindText, IsFindError := Selector.TextContent(); IsFindError == nil {
			card.Type = strings.TrimSpace(FindText)
		}
	}

	// Статус дела
	if Selector, _ := core.page.QuerySelector(`div[class=b-case-header-desc]`); Selector != nil { // Если найден такой блок
		// Берём текстовое значение и проверяем его на ошибку
		if FindText, IsFindError := Selector.TextContent(); IsFindError == nil {
			card.Status = strings.TrimSpace(FindText)
		}
	}

	// Следующее заседание
	if Selector, _ := core.page.QuerySelector(`div[class=b-instanceAdditional]`); Selector != nil { // Если найден такой блок
		// Берём текстовое значение и проверяем его на ошибку
		if FindText, IsFindError := Selector.TextContent(); IsFindError == nil {
			// FindText = strings.TrimSpace(FindText)
			FindText = strings.ReplaceAll(FindText, "Следующее заседание:", "")
			FindText = strings.TrimSpace(FindText)
			strs := strings.Split(FindText, ",")
			if len(strs) == 3 {
				fmt.Println(">" + strs[0] + " " + strs[1] + "< - >" + strs[2] + "<")
				// Локация заседания
				card.Next.Location = strings.TrimSpace(strs[2])

				// Время сделающего заседания
				ParseTime, ErrorParse := time.Parse("02.01.2006 15:04", strings.TrimSpace(strs[0])+" "+strings.TrimSpace(strs[1]))
				if ErrorParse == nil {
					card.Next.Date = ParseTime
				}
			} else {
				fmt.Println("Не могу преобразовать данные для следующего заседания. Нужно добавить обработчик.\n" + FindText)
			}
		}
	}

	// Сперва пропарсим главные значения карточек
	MainsH, err := core.page.QuerySelectorAll(`div[class="b-chrono-item-header js-chrono-item-header page-break"] div`)
	if err != nil {
		return card, err // could not get entries
	}
	if len(MainsH) != 0 { // Если ненулевое к-во элементов
		// Выделяем память для карточек
		card.Slips = make([]struct {
			Main  HistoryMain
			Slave []HistorySlave
		}, len(MainsH))
		for IndexMain, mainH := range MainsH { // Парсим каждую главную карточку
			// Название инстанции суда
			if Selector, _ := mainH.QuerySelector(`div[class=l-col] strong`); Selector != nil { // Если найден такой блок
				// Берём текстовое значение и проверяем его на ошибку
				if FindText, IsFindError := Selector.TextContent(); IsFindError == nil {
					FindText = strings.TrimSpace(FindText)
					card.Slips[IndexMain].Main.InstanceName = FindText
				}
			}

			// Дата
			if Selector, _ := mainH.QuerySelector(`div[class=l-col] span[class=b-reg-date]`); Selector != nil { // Если найден такой блок
				// Берём текстовое значение и проверяем его на ошибку
				if FindText, IsFindError := Selector.TextContent(); IsFindError == nil {
					FindText = strings.TrimSpace(FindText)
					ParseTime, ErrorParse := time.Parse("02.01.2006", FindText)
					if ErrorParse == nil {
						card.Slips[IndexMain].Main.Date = ParseTime
					}
				}
			}

			// Инстанции
			if Selector, _ := mainH.QuerySelector(`div[class=l-col] span[class=b-reg-incoming_num]`); Selector != nil { // Если найден такой блок
				// Берём текстовое значение и проверяем его на ошибку
				if FindText, IsFindError := Selector.TextContent(); IsFindError == nil {
					FindText = strings.TrimSpace(FindText)
					card.Slips[IndexMain].Main.Number = FindText
				}
			}

			// Отчет по датам публикаций
			if Selector, _ := mainH.QuerySelector(`div[class=r-col] h4 span[class=b-indentIcon] a`); Selector != nil { // Если найден такой блок
				// Берём текстовое значение и проверяем его на ошибку
				if FindText, EroorFind := Selector.GetAttribute("href"); EroorFind == nil {
					card.Slips[IndexMain].Main.UrlReport = FindText
				}
			}

			// Номер инстанции
			if Selector, _ := mainH.QuerySelector(`div[class=r-col] h4 span strong[class=b-case-instance-number]`); Selector != nil { // Если найден такой блок
				// Берём текстовое значение и проверяем его на ошибку
				if FindText, IsFindError := Selector.TextContent(); IsFindError == nil {
					FindText = strings.TrimSpace(FindText)
					card.Slips[IndexMain].Main.NumberInstance = FindText
				}
			}

			// Название суда
			if Selector, _ := mainH.QuerySelector(`div[class=r-col] h4 span span[class=instantion-name] a`); Selector != nil { // Если найден такой блок
				// Берём текстовое значение и проверяем его на ошибку
				if FindText, IsFindError := Selector.TextContent(); IsFindError == nil {
					FindText = strings.TrimSpace(FindText)
					card.Slips[IndexMain].Main.Cour = FindText
				}
			}

			// Ссылка на суд
			if Selector, _ := mainH.QuerySelector(`div[class=r-col] h4 span span[class=instantion-name] a`); Selector != nil { // Если найден такой блок
				// Берём текстовое значение и проверяем его на ошибку
				if FindText, EroorFind := Selector.GetAttribute("href"); EroorFind == nil {
					card.Slips[IndexMain].Main.UrlCour = FindText
				}
			}
		}
	}

	return card, ErrorParse
}
