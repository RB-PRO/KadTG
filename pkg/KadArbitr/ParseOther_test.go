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

	// --- 1 test ---
	// https://kad.arbitr.ru/Card/72197155-c243-47d3-b328-2c421391754a
	tests[0].Answer.Slips = make([]struct { // Выделяем память
		Main  KadArbitr.HistoryMain
		Slave []KadArbitr.HistorySlave
	}, 2)
	tests[0].Answer.Slips[0].Main = KadArbitr.HistoryMain{
		InstanceName:   "Апелляционная инстанция",
		Date:           time.Date(2015, time.December, 2, 0, 0, 0, 0, time.UTC),
		Number:         "21АП-1781/15 (1)",
		UrlReport:      "https://kad.arbitr.ru/PublishReport?instanceId=7265debb-0a04-4d27-b407-1ee987eeab99&caseId=72197155-c243-47d3-b328-2c421391754a",
		NumberInstance: "21АП-1781/2015",
		Cour:           "21 арбитражный апелляционный суд",
		UrlCour:        "http://21aas.arbitr.ru",
		FileName:       "Оставить решение суда без изменения, жалобу без удовлетворения",
		FileLink:       "https://kad.arbitr.ru/Kad/PdfDocument/72197155-c243-47d3-b328-2c421391754a/b95dbc3e-bb63-4cee-9db3-b97cc3570dc0/A84-610-2015_20151202_Reshenija_i_postanovlenija.pdf",
	}
	tests[0].Answer.Slips[1].Main = KadArbitr.HistoryMain{
		InstanceName:   "Первая инстанция",
		Date:           time.Date(2015, time.August, 17, 0, 0, 0, 0, time.UTC),
		UrlReport:      "https://kad.arbitr.ru/PublishReport?instanceId=3a42aeb0-0a2e-4102-83aa-9f71f6a54911&caseId=72197155-c243-47d3-b328-2c421391754a",
		NumberInstance: "А84-610/2015",
		Cour:           "АС города Севастополя",
		UrlCour:        "http://sevastopol.arbitr.ru",
		FileName:       "В иске отказать полностью",
		FileLink:       "https://kad.arbitr.ru/Kad/PdfDocument/72197155-c243-47d3-b328-2c421391754a/fd508d2e-23cc-4370-8ba4-f8b705005d90/A84-610-2015_20150817_Reshenija_i_postanovlenija.pdf",
	}

	// --- 2 test ---
	// https://kad.arbitr.ru/Card/7691c97c-01ce-4104-b00d-5125a27a44fc
	tests[1].Answer.Slips = make([]struct { // Выделяем память
		Main  KadArbitr.HistoryMain
		Slave []KadArbitr.HistorySlave
	}, 1)
	tests[1].Answer.Slips[0].Main = KadArbitr.HistoryMain{
		InstanceName:   "Первая инстанция",
		UrlReport:      "https://kad.arbitr.ru/PublishReport?instanceId=a9a197b1-2b37-431c-bd07-5b4734375ebe&caseId=7691c97c-01ce-4104-b00d-5125a27a44fc",
		NumberInstance: "А40-84224/2023",
		Cour:           "АС города Москвы",
		UrlCour:        "http://msk.arbitr.ru",
	}

	// --- 3 test ---
	// https://kad.arbitr.ru/Card/b9b800bb-eb4d-4826-87c8-d3259bc822de
	tests[2].Answer.Slips = make([]struct { // Выделяем память
		Main  KadArbitr.HistoryMain
		Slave []KadArbitr.HistorySlave
	}, 1)
	tests[2].Answer.Slips[0].Main = KadArbitr.HistoryMain{
		InstanceName:   "Первая инстанция",
		Date:           time.Date(2023, time.February, 21, 0, 0, 0, 0, time.UTC),
		UrlReport:      "https://kad.arbitr.ru/PublishReport?instanceId=4fc3ac52-398c-4b63-878b-326f395b8b7a&caseId=b9b800bb-eb4d-4826-87c8-d3259bc822de",
		NumberInstance: "А84-8614/2022",
		Cour:           "АС города Севастополя",
		UrlCour:        "http://sevastopol.arbitr.ru",
		FileName:       "Иск удовлетворить полностью",
		FileLink:       "https://kad.arbitr.ru/Kad/PdfDocument/b9b800bb-eb4d-4826-87c8-d3259bc822de/28cd1c8e-d5e4-4162-89cb-80614adab524/A84-8614-2022_20230221_Reshenija_i_postanovlenija.pdf",
	}

	// Создаём ядро
	core, ErrorCore := KadArbitr.NewCore()
	if ErrorCore != nil {
		t.Error(ErrorCore)
	}

	// Цикл по тестовым парам
	for _, tt := range tests {
		fmt.Println("Ссылка:", tt.url)
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

		// Цикл по всем карточкам
		for index, value := range tt.Answer.Slips {
			// Суд
			if value.Main.Cour != card.Slips[index].Main.Cour {
				t.Errorf(`Slips[index].Main.Cour another. Вместо "%v", получено "%v".`, value.Main.Cour, card.Slips[index].Main.Cour)
			}

			// Инстанция суда
			if value.Main.InstanceName != card.Slips[index].Main.InstanceName {
				t.Errorf(`Slips[index].Main.InstanceName another. Вместо "%v", получено "%v".`, value.Main.InstanceName, card.Slips[index].Main.InstanceName)
			}

			// Дата
			if value.Main.Date.GoString() != card.Slips[index].Main.Date.GoString() {
				t.Errorf(`Slips[index].Main.Number another. Вместо "%v", получено "%v".`, value.Main.Date.GoString(), card.Slips[index].Main.Date.GoString())
			}

			// Номер дела
			if value.Main.Number != card.Slips[index].Main.Number {
				t.Errorf(`Slips[index].Main.Number another. Вместо "%v", получено "%v".`, value.Main.Number, card.Slips[index].Main.Number)
			}

			// Отчет по датам публикаций
			if value.Main.UrlReport != card.Slips[index].Main.UrlReport {
				t.Errorf(`Slips[index].Main.UrlReport another. Вместо "%v", получено "%v".`, value.Main.UrlReport, card.Slips[index].Main.UrlReport)
			}

			// Номер инстанции
			if value.Main.NumberInstance != card.Slips[index].Main.NumberInstance {
				t.Errorf(`Slips[index].Main.NumberInstance another. Вместо "%v", получено "%v".`, value.Main.NumberInstance, card.Slips[index].Main.NumberInstance)
			}

			// Ссылка на суд
			if value.Main.UrlCour != card.Slips[index].Main.UrlCour {
				t.Errorf(`Slips[index].Main.UrlCour another. Вместо "%v", получено "%v".`, value.Main.UrlCour, card.Slips[index].Main.UrlCour)
			}

			// Название файла
			if value.Main.FileName != card.Slips[index].Main.FileName {
				t.Errorf(`Slips[index].Main.FileName another. Вместо "%v", получено "%v".`, value.Main.FileName, card.Slips[index].Main.FileName)
			}

			// Ссылка на файл
			if value.Main.FileLink != card.Slips[index].Main.FileLink {
				t.Errorf(`Slips[index].Main.FileLink another. Вместо "%v", получено "%v".`, value.Main.FileLink, card.Slips[index].Main.FileLink)
			}
		}

	}
}
