package KadArbitr

import "time"

const ParticipantAll string = "-1"       // Любой
const ParticipantPlaintiff string = "0"  // Истец
const ParticipantRespondent string = "1" // Ответчик
const ParticipantThird string = "2"      // Третье лицо
const ParticipantAnother string = "3"    // Иное лицо

// Структура запросов
type Request struct {
	// * Участник дела *
	Participant []struct {
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

	// * Судья *
	Judgs []string

	// * Суд *
	Court []string

	// * Номер дела *
	Number []string

	// * Дата регистрации дела *
	DateTo   time.Time // Дата начала
	DateFrom time.Time // Дата конца

}
