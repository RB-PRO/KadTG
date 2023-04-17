package KadArbitr

import "time"

const (
	ParticipantAll        string = ""  // Любой - Категория участника дела
	ParticipantPlaintiff  string = "0" // Истец - Категория участника дела
	ParticipantRespondent string = "1" // Ответчик - Категория участника дела
	ParticipantThird      string = "2" // Третье лицо - Категория участника дела
	ParticipantAnother    string = "3" // Иное лицо - Категория участника дела
)

const (
	ModeSearch         string = ""               // Найти. Обычный поиск. - Параметры поиска
	ModeAdministrative string = "administrative" // Административные - Параметры поиска
	ModeCivil          string = "civil"          // Гражданские - Параметры поиска
	ModeBankruptcy     string = "bankruptcy"     // Банкротные - Параметры поиска
)

// Структура запросов
type Request struct {

	// Участник дела + настройка стороны
	Part []Participant

	// Судья + инстанция
	Judg []Judgs

	// Суд
	Court []string

	// Номер дела
	Number []string

	// Дата регистрации дела
	DateFrom time.Time // ОТ
	DateTo   time.Time // ДО

	// Судебные поручения
	// true эквивалентно судебным поречениям
	LegendCheckbox bool

	// Параметры поиска
	//	-"a" - Административные
	//	-"c" - Гражданские
	//	-"b" - Банкротные
	//	-"" - Найти
	SearchCases string
}

// Структура Участника дела
type Participant struct {
	// Значение запроса:
	//	название, ИНН или ОГРН
	Value string

	// Настрйока запроса:
	//	- ParticipantAll - Любой
	//	- ParticipantPlaintiff - Истец
	//	- ParticipantRespondent - Ответчик
	//	- ParticipantThird - Третье лицо
	//	- ParticipantAnother - Иное лицо
	Settings string
}

// Структура Судьи вместе с инстанцией
type Judgs struct {
	Value    string // Судья
	Instance string // Инстанция
}
