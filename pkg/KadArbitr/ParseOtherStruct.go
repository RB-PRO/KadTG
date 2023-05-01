package KadArbitr

import "time"

// Структура для парсинга содержимого каждого дела.
type Card struct {

	// ID данного судебного дела. Используется для формирования запросов
	CaseID string

	// Сумма исковых требований
	Coast int

	// Типо судебного дела
	Type string

	// Рассмотрение дела завершено
	Status string

	// Следующее заседание
	Next struct {
		Date     time.Time // Дата проведения следующего заседания
		Location string    // Локация
	}

	Slips []struct {

		// В cURL запросе является полем id. Является ID карточки
		DataID string

		// Главная карточка
		Main HistoryMain

		// Подкарточки
		Slave []HistorySlave
	}
}

// Главная карточка. То что находится во главе этих карт судов. Похоже на обычную карточку истории за исключением дополнительных полей.
type HistoryMain struct {
	// --- Левая колонка ---
	// Инстанция суда
	InstanceName string

	// Дата
	Date time.Time

	// Номер дела
	Number string

	// --- Правая колонка ---
	// Отчет по датам публикаций
	UrlReport string

	// Номер инстанции
	NumberInstance string

	// Суд
	Cour string

	// Ссылка на суд
	UrlCour string

	// Название файла
	FileName string

	// Ссылка на файл
	FileLink string
}

// Карточки
// Это структура одного дела из карточки судебных событий
// По сути это история движения дела со следующими параметрами:
//
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
