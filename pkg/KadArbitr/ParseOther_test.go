package KadArbitr_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/RB-PRO/KadTG/pkg/KadArbitr"
)

func TestParseCard(t *testing.T) {
	// Табличные тесты
	times := make([]time.Time, 3)
	times[1], _ = time.Parse("02.01.2006 15:04", "29.06.2023 16:05") // "Следующее заседание: 29.06.2023, 16:05 , Зал судебных заседаний № 10063"
	fmt.Println(times[1])
	var tests = []struct {
		url    string
		Answer KadArbitr.Card
	}{
		{
			url: "https://kad.arbitr.ru/Card/72197155-c243-47d3-b328-2c421391754a",
			Answer: KadArbitr.Card{
				Type:   "экономические споры по гражданским правоотношениям",
				Status: "Рассмотрение дела завершено",
				Coast:  0,
			},
		},
		{
			url: "https://kad.arbitr.ru/Card/7691c97c-01ce-4104-b00d-5125a27a44fc",
			Answer: KadArbitr.Card{
				Type:   "о несостоятельности (банкротстве) организаций и граждан",
				Status: "Рассматривается в первой инстанции",
				Coast:  0,
				Next: struct {
					Date     time.Time
					Location string
				}{
					Date:     times[1],
					Location: "Зал судебных заседаний № 10063",
				},
			},
		},
		{
			url: "https://kad.arbitr.ru/Card/b9b800bb-eb4d-4826-87c8-d3259bc822de",
			Answer: KadArbitr.Card{
				Type:   "экономические споры по гражданским правоотношениям",
				Status: "Рассмотрение дела завершено",
				Coast:  60000,
			},
		},
		// https://kad.arbitr.ru/Card/0c433692-2f74-43b3-8f37-4e1602ca4d93 - тут мало заполнены поля

		// https://kad.arbitr.ru/Card/e219ee0c-4eea-459a-951c-210d7e203975 - Тут есть подпись
	}
	// tests[0].Answer.Slips = append(tests[0].Answer.Slips, []struct{Main KadArbitr.HistoryMain; Slave []KadArbitr.HistorySlave})

	// Выделяем память
	tests[0].Answer.Slips = make([]struct {
		Main  KadArbitr.HistoryMain
		Slave []KadArbitr.HistorySlave
	}, 2)
	tests[0].Answer.Slips[0].Main = KadArbitr.HistoryMain{
		InstanceName: "Апелляционная инстанция",
		// Date time.Time
		Number:         "21АП-1781/15 (1)",
		UrlReport:      "",
		NumberInstance: "",
		Cour:           "",
		UrlCour:        "",
		FileName:       "",
		FileLink:       "",
	}

	// Создаём ядро
	core, ErrorCore := KadArbitr.NewCore()
	if ErrorCore != nil {
		t.Error(ErrorCore)
	}

	// Цикл по тестовым парам
	for _, tt := range tests {
		fmt.Println("Ссылка", tt.url)
		card, ErrorCard := core.ParseCard(tt.url)
		if ErrorCard != nil {
			t.Error(ErrorCard)
		}

		// Статус дела
		if tt.Answer.Status != card.Status {
			t.Errorf(`Status another. Вместо "%v", получено "%v".`, tt.Answer.Status, card.Status)
		}

		// Тип дела
		if tt.Answer.Type != card.Type {
			t.Errorf(`Type another. Вместо "%v", получено "%v".`, tt.Answer.Type, card.Type)
		}

		// Следующее судебное заседание, локация
		if tt.Answer.Next.Location != card.Next.Location {
			t.Errorf(`Next.Location another. Вместо "%v", получено "%v".`, tt.Answer.Next.Location, card.Next.Location)
		}
		if tt.Answer.Next.Date != card.Next.Date {
			t.Errorf(`Next.Date another. Вместо "%v", получено "%v".`, tt.Answer.Next.Date, card.Next.Date)
		}
	}
}
