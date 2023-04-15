package KadArbitr

// Парсинг всех страниц документа с учётом колличества страниц
// На вход структуру запроса, а на выходе массив полученных данных

// Структура парсинга всех страниц
type Parse struct {
	Settings PageSettings // Настройки запроса
	Data     []Data       // Структура результатов парсинга
	Request  Request      // Структура запросов
}

// Структура "настроек" для запроса
type PageSettings struct {
	DocumentsPageSize   int // К-во записей на одной странице
	DocumentsPage       int // Текущая страница
	DocumentsTotalCount int // Всего записей по запросу
	DocumentsPagesCount int // Всего страниц по запросу
}

func (core *CoreReq) ParseAll() {

}
