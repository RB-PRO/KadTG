package KadArbitr

import (
	"testing"
	"time"
)

func TestNumberTotalPages(t *testing.T) {
	var tests = []struct {
		in  int
		out int
	}{
		{0, 0},
		{1, 1},
		{584, 24},
		{27259128, 40},
	}

	for _, e := range tests {
		answer := NumberTotalPages(e.in)
		if answer != e.out {
			t.Errorf("Для %v. Результат %v, а должно быть %v.", e.in, answer, e.out)
		}
	}
}
func TestNextPage(t *testing.T) {
	// Создаём ядро
	core, ErrorCore := NewCore()
	if ErrorCore != nil {
		t.Error(ErrorCore)
	}
	core.Screen("screens/Next1.jpg")
	req := Req2() // Создаём запрос на поиск
	// Заполнение формы поиска
	ErrorReq := core.FillReqestOne(req)
	if ErrorReq != nil {
		t.Error(ErrorReq)
	}
	core.Screen("screens/Next2.jpg")

	ErrorSearch := core.Search(req)
	if ErrorSearch != nil {
		t.Error(ErrorSearch)
	}
	core.Screen("screens/Next3.jpg")

	ErrorNext := core.NextPage()
	if ErrorNext != nil {
		t.Error(ErrorNext)
	}
	time.Sleep(2 * time.Second)
	core.Screen("screens/Next4.jpg")
}

// Получить тестовый запрос
// В этом запросе дохуя(боле 4600) записей
func Req2() Request {
	return Request{
		// Стороны
		Part: []Participant{
			{
				Value:    `7736050003`,   // Истец
				Settings: ParticipantAll, // Категория истца
			},
		},
	}
}
