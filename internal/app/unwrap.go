package app

import (
	"strings"
	"time"

	"github.com/RB-PRO/KadTG/pkg/KadArbitr"
)

// Преобразовать входную строку в структуру запроса
func unwrap(input string) (req KadArbitr.Request, err error) {
	req.Judg = []KadArbitr.Judgs{}
	req.Part = []KadArbitr.Participant{}

	strs := strings.Split(input, "\n") // Получить строки
	for _, str := range strs {
		f := string(str[0])

		// Участник дела
		if strings.Contains(f, "1") {
			str = strings.ReplaceAll(str, "1. ", "")
			strss := strings.Split(str, ";")
			req.Part = append(req.Part, KadArbitr.Participant{
				Value:    strings.TrimSpace(strss[0]),
				Settings: strings.TrimSpace(strss[1]),
			})
		}

		// Судья
		if strings.Contains(f, "2") {
			str = strings.ReplaceAll(str, "2. ", "")
			strss := strings.Split(str, ";")
			req.Judg = append(req.Judg, KadArbitr.Judgs{
				Value:    strings.TrimSpace(strss[0]),
				Instance: strings.TrimSpace(strss[1]),
			})
		}

		// Номер дела
		if strings.Contains(f, "3") {
			str = strings.ReplaceAll(str, "3. ", "")
			req.Number = append(req.Number, strings.TrimSpace(str))
		}

		// Суд
		if strings.Contains(f, "4") {
			str = strings.ReplaceAll(str, "4. ", "")
			req.Court = append(req.Number, strings.TrimSpace(str))
		}

		// Дата регистрации С
		if strings.Contains(f, "5") {
			str = strings.ReplaceAll(str, "5. ", "")
			if times, err := time.Parse("02.01.2006", strings.TrimSpace(str)); err == nil {
				req.DateFrom = times
			}
		}

		// Дата регистрации ДО
		if strings.Contains(f, "6") {
			str = strings.ReplaceAll(str, "6. ", "")
			if times, err := time.Parse("02.01.2006", strings.TrimSpace(str)); err == nil {
				req.DateTo = times
			}
		}

		// Параметр поиска
		if strings.Contains(f, "7") {
			str = strings.ReplaceAll(str, "7. ", "")
			req.SearchCases = strings.TrimSpace(str)
			switch strings.TrimSpace(str) {
			case "a":
				req.SearchCases = KadArbitr.ModeAdministrative
			case "c":
				req.SearchCases = KadArbitr.ModeCivil
			case "b":
				req.SearchCases = KadArbitr.ModeBankruptcy
			default:
				req.SearchCases = KadArbitr.ModeSearch
			}
		}
	}
	return req, nil
}

/*
1. [ИНН или компания]; [сторона( "0" - Истец, "1" - Ответчик,"2" - Третье лицо, "3" - Иное лицо)]
2. [судья]; [инстанция]
3. [номер дела]
4. [Дата регистрации С]
5. [Дата регистрации ДО]
6. [Параметр поиска("a" - Административные,"c" - Гражданские, "b" - Банкротные, "o" - Найти обычным поиском)]
*/

///////////////////
/*
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

	//////////////////////////////////////

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

*/
