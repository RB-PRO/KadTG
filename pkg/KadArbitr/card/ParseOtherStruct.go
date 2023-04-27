package card

import "time"

// Структура для парсинга содержимого каждого дела.
type Carda struct {
	// Сумма исковых требований
	Coast int

	// Рассмотрение дела завершено
	Status string

	// Следующее заседание
	Next struct {
		Date     time.Time // Дата проведения следующего заседания
		Location string    // Локация
	}

	Slips []struct {
		// Главная карточка
		Main HistoryMain

		// Подкарточки
		Slave []HistorySlave
	}
}

// Главная карточка. То что находится во главе этих карт судов. Похоже на обычную карточку истории за исключением дополнительных полей.
type HistoryMain struct {
	// Инстанция суда
	InstanceName string

	// Дата
	Date time.Time

	// Номер дела
	Number string

	// Инстанция суда
	Instance struct {
		// Отчет по датам публикаций
		UrlReport string

		// Номер инстанции
		NumberInstance string

		// Суд
		Cour string

		// Ссылка на суд
		UrlCour string
	}

	// Приложение. Либо файл, либо название
	Application struct {
		// Название файла
		Name string

		// Ссылка на файл
		Link string
	}
}

// Карточки
// Это структура одного дела из карточки судебных событий
// По сути это история движения дела со следующими параметрами:
//	-
//	-
//	-
//	-
type HistorySlave struct {
	// Дата дела
	Date time.Time

	// Тип дела
	Type string

	// Информация о деле(там содержится информация о месте встрече и сумме иска)
	Info string

	// Поле в котором могут быть записаны судьи или суд
	JudgeOrCourt []string

	// Приложение. Либо файл, либо название
	Application struct {
		// Название файла
		Name string

		// Ссылка на файл
		Link string

		// Судебный состав
		JudicialComposition []string

		// Судья-докладчик
		JudgeSpeaker []string

		// Судьи
		Judges []string
	}

	// Дата публикации
	DatePost struct {
		// Дата+время публикации
		time.Time

		// Ссылка на какой-то документ с датой
		URL string
	}
}
