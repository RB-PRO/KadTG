package KadArbitr

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
		return pr, nil
	}

	// В случае, если найдена только одна страница
	// то парсим и выводим результат
	if pr.Settings.DocumentsPagesCount == 1 {
		pr.Data, ErrorAll = core.Parse()
		if ErrorAll != nil {
			return pr, ErrorAll
		}
		return pr, nil
	}

	//

	return pr, ErrorAll
}
