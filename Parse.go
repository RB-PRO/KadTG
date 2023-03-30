package KadArbitr

import (
	"fmt"
	"time"
)

type Data struct {
	// 1 Колонка
	Date      time.Time // Дата
	Number    string    // Номер дела
	UrlNumber string    // Ссылка на дело

	// 2 Колонка
	Judge    string // Судья
	Instance string // Инстанция

	// 3 Колонка
	Plaintiff Side // Истец

	// 4 Колонка
	Respondent Side // Ответчик
}

// Сторона конфликта
type Side struct {
	Name   string // Название компании
	Adress string // Адресс
	INN    string // ИНН
}

// b-cases
func (core *CoreReq) Parse() ([]Data, error) {
	Datas := make([]Data, 0)

	entries, err := core.page.QuerySelectorAll("table[class='b-cases'] > tbody > tr")
	if err != nil {
		return nil, err // could not get entries
	}
	for i, entry := range entries {
		var AppendData Data

		fmt.Println(i)

		// 1 Колонка
		num, errorNum := entry.QuerySelector("td.num")
		if errorNum == nil {

			// Дата
			DateSelector, DateError := num.QuerySelector("span")
			if DateError == nil {
				DateText, FindDateText := DateSelector.InnerText()
				if FindDateText == nil {
					TimeDate, FindTime := time.Parse("02.01.2006", DateText)
					if FindTime == nil {
						AppendData.Date = TimeDate
					}
				}
			}

			// Номер дела + ссылка на дело
			NumberSelector, NumberError := num.QuerySelector("a")
			if NumberError == nil {
				// Номер дела
				NumberText, FindNumberText := NumberSelector.InnerText()
				if FindNumberText == nil {
					AppendData.Number = NumberText
				}

				// Ссылка на дело
				HrefStr, IsHref := NumberSelector.GetAttribute("href")
				if IsHref == nil {
					AppendData.UrlNumber = HrefStr
				}
			}

		}

		// Добавляем новую структуру в выходной массив структур
		Datas = append(Datas, AppendData)
	}

	return Datas, nil
}
