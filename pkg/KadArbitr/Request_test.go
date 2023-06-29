package KadArbitr

import (
	"fmt"
	"io/ioutil"
	"testing"
	"time"
)

func TestFillReqest(t *testing.T) {
	// Пытаюсь достучаться до дела
	// https://kad.arbitr.ru/Card/23fee179-ae11-4036-9139-bae0babcaea7

	// Создаём ядро
	core, ErrorCore := NewCore()
	if ErrorCore != nil {
		t.Error(ErrorCore)
	}

	core.Screen("screens/Req1.jpg")

	req := Req1() // Создаём запрос на поиск

	// Заполнение формы поиска
	ErrorReq := core.FillReqestOne(req)
	if ErrorReq != nil {
		t.Error(ErrorReq)
	}

	core.Screen("screens/Req2.jpg")

	// save
	html, _ := core.page.QuerySelector("body")
	htmlB, _ := html.InnerHTML()
	ioutil.WriteFile("output.html", []byte(htmlB), 0644)

	core.Search(req)
	core.Screen("screens/Req3.jpg")

	data, _ := core.Parse()
	fmt.Println("len", len(data))
	fmt.Printf("%+v\n", data[0])
	fmt.Printf("суд %+v\n", data[0].Instance)

	core.Screen("screens/Req4.jpg")
}

// Получить тестовый запрос
// В этом запросе всего одна запись
func Req1() Request {
	return Request{
		// Стороны
		Part: []Participant{
			{
				Value:    `ООО М4 Б2Б МАРКЕТПЛЕЙС`, // Истец
				Settings: ParticipantPlaintiff,     // Категория истца
			},
			{
				Value:    `Деева Екатерина Николаевна`, // Истец
				Settings: ParticipantRespondent,        // Категория ответчика
			},
		},

		// Судья
		Judg: []Judgs{
			{
				Value:    `Снегур А. А.`,
				Instance: `Суд по интеллектуальным правам`,
			},
		},

		// Суд
		Court: []string{`Суд по интеллектуальным правам`},
		//Court: []string{`Верховный Суд РФ`},

		// Номер дела
		Number: []string{`СИП-344/2023`},

		// Дата регистрации
		DateTo:   time.Date(2023, time.April, 14, 0, 0, 0, 0, time.Local),
		DateFrom: time.Date(2023, time.April, 14, 0, 0, 0, 0, time.Local),

		// Судебные поручения
		LegendCheckbox: false,

		// Параметры поиска
		SearchCases: ModeCivil,
	}
}
