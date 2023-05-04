package KadArbitr

import "fmt"

// Парсинг всех страниц документа с учётом колличества страниц
// На вход структуру запроса, а на выходе массив полученных данных

// Структура парсинга всех страниц
type Parse struct {
	Settings PageSettings // Настройки запроса
	Data     []Data       // Структура результатов парсинга
	//Request  Request      // Структура запросов
}

// Структура "настроек" для запроса
type PageSettings struct {
	DocumentsPageSize   int // К-во записей на одной странице
	DocumentsPage       int // Текущая страница
	DocumentsTotalCount int // Всего записей по запросу
	DocumentsPagesCount int // Всего страниц по запросу
}

// Спарсить все страницы по запросу
func (core *CoreReq) ParseAll() (pr Parse, ErrorAll error) {
	// pr - Объект парсинга всех страниц сразу.
	// Условие - Должна быть открыта исходная страница https://kad.arbitr.ru/

	// Получаем настройки при первом открытии страницы
	pr.Settings, ErrorAll = core.Settings()
	if ErrorAll != nil {
		return Parse{}, ErrorAll
	}

	// Если найдено ноль записей, то и парсить соответственно нечего
	if pr.Settings.DocumentsTotalCount == 0 {
		return Parse{}, nil
	}

	// В случае, если найдена только одна страница
	// то парсим и выводим результат
	if pr.Settings.DocumentsPagesCount == 1 {
		pr.Data, ErrorAll = core.Parse()
		return pr, ErrorAll
	}

	// Если страниц больше 1
	if pr.Settings.DocumentsPagesCount > 1 {

		// Парсим текущую страницу
		pr.Data, ErrorAll = core.Parse()
		if ErrorAll != nil {
			return Parse{}, ErrorAll
		}

		// Цикл по всем страницам
		for pr.Settings.DocumentsPage = 2; pr.Settings.DocumentsPage <= pr.Settings.DocumentsPagesCount; pr.Settings.DocumentsPage++ {
			fmt.Println(pr.Settings.DocumentsPage, "из", pr.Settings.DocumentsPagesCount)

			// Массив записей на странице
			var data []Data

			// Следующая страница
			ErrorNext := core.NextPage()
			if ErrorNext != nil {
				fmt.Println("ParseAll:", ErrorNext)
				//return Parse{}, ErrorNext
			}

			// Парсим страницы
			data, ErrorAll = core.Parse()
			if ErrorAll != nil {
				return Parse{}, ErrorAll
			}

			// Добавляем в готвоый массив
			pr.Data = append(pr.Data, data...)
		}
	}

	return pr, ErrorAll
}
