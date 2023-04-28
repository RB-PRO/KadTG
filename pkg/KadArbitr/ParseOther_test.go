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
	times[1], _ = time.Parse("02.01.2006, 15:04", "Следующее заседание: 29.06.2023, 16:05 , Зал судебных заседаний № 10063")
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
					Date:     times[0],
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

	fmt.Println(tests)

	// Создаём ядро
	core, ErrorCore := KadArbitr.NewCore()
	if ErrorCore != nil {
		t.Error(ErrorCore)
	}
	core.Screen("screens/Card1.jpg")

	card, ErrorCard := core.ParseCard("https://kad.arbitr.ru/Card/72197155-c243-47d3-b328-2c421391754a")
	if ErrorCard != nil {
		t.Error(ErrorCard)
	}
	core.Screen("screens/Card4.jpg")

	fmt.Printf("%+v", card)

}
