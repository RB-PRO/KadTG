package KadArbitr

import "strconv"

// Получить структуру настроек со страницы данных
func (core *CoreReq) Settings() (Settings PageSettings, ErrorFind error) {
	var Value int

	// Количество записей на одной странице
	Value, ErrorFind = core.settingsOne("documentsPageSize")
	if ErrorFind != nil {
		return Settings, ErrorFind
	}
	Settings.DocumentsPageSize = Value

	// Текущая страница
	Value, ErrorFind = core.settingsOne("documentsPage")
	if ErrorFind != nil {
		return Settings, ErrorFind
	}
	Settings.DocumentsPage = Value

	// Количество записей
	Value, ErrorFind = core.settingsOne("documentsTotalCount")
	if ErrorFind != nil {
		return Settings, ErrorFind
	}
	Settings.DocumentsTotalCount = Value

	// Количество страниц
	Value, ErrorFind = core.settingsOne("documentsPagesCount")
	if ErrorFind != nil {
		return Settings, ErrorFind
	}
	Settings.DocumentsPagesCount = Value

	return Settings, nil
}

// Получить данные одной настройки из input и вернуть значение в int+error
func (core *CoreReq) settingsOne(ID string) (Value int, ErrorFind error) {

	// Собираем настройки данной страницы
	Selector, ErrorFind := core.page.QuerySelector("input[id=" + ID + "]")
	if ErrorFind != nil {
		return 0, ErrorFind
	}

	// Получаем данные по атрибуту
	SelectorInnerHTML, ErrorFind := Selector.GetAttribute("value")
	if ErrorFind != nil {
		return 0, ErrorFind
	}

	if SelectorInnerHTML == "" {
		SelectorInnerHTML = "0"
	}

	// Переводим string в int
	if Value, ErrorFind = strconv.Atoi(SelectorInnerHTML); ErrorFind != nil {
		return 0, ErrorFind
	}

	return Value, nil
}
