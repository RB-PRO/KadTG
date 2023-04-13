package KadArbitr

import (
	"fmt"
	"strings"
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

	// ***

	entries, err := core.page.QuerySelectorAll("table[class='b-cases'] > tbody > tr")
	if err != nil {
		return nil, err // could not get entries
	}
	for _, TR := range entries {
		var AppendData Data

		// Массив колонок. Их должно быть 4 штуки
		TD, ErrorTd := TR.QuerySelectorAll("td")
		if ErrorTd != nil {
			continue
		}

		// ************* //
		// * 1 колонка * //
		// ************* //
		// Дата
		DateSelector, DateError := TD[0].QuerySelector("span")
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
		NumberSelector, NumberError := TD[0].QuerySelector("a")
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

		// ************* //
		// * 2 колонка * //
		// ************* //
		// Судья
		Judge, ErrorJudge := TD[1].QuerySelector("div div.judge")
		if ErrorJudge == nil {
			JudgeText, IsJudge := Judge.InnerText()
			if IsJudge == nil {
				AppendData.Judge = JudgeText
			}
		}

		// // Инстанция
		TecalCourt, ErrorCourt := TD[1].QuerySelector("div div:last-of-type")
		if ErrorCourt == nil {
			AppendData.Instance, _ = TecalCourt.InnerText()
		}

		// ************* //
		// * 3 колонка * //
		// ************* //
		AppendData.Plaintiff = Side{} // Определить Истца
		Plaintiff, ErrorPlaintiff := TD[2].QuerySelector("span[class=js-rolloverHtml]")
		if ErrorPlaintiff == nil {

			// Название компании
			Name, ErrorName := Plaintiff.QuerySelector("strong")
			if ErrorName == nil {
				AppendData.Plaintiff.Name, _ = Name.InnerText()
			}

			// Адрес
			// AppendData.Plaintiff.Adress, _ = Plaintiff.InnerText()

			// ИНН
			INN, ErrorINN := Plaintiff.QuerySelector("div")
			if ErrorINN == nil {
				AppendData.Plaintiff.INN, _ = INN.InnerText()
				AppendData.Plaintiff.INN = strings.ReplaceAll(AppendData.Plaintiff.INN, "ИНН:", "")
				AppendData.Plaintiff.INN = strings.TrimSpace(AppendData.Plaintiff.INN)
			}
		}

		// ************* //
		// * 4 колонка * //
		// ************* //
		AppendData.Respondent = Side{} // Определить Ответчика
		Respondent, ErrorRespondent := TD[3].QuerySelector("span[class=js-rolloverHtml]")
		if ErrorRespondent == nil {

			// Название компании
			Name, ErrorName := Respondent.QuerySelector("strong")
			if ErrorName == nil {
				if Value, FindError := Name.InnerText(); FindError != nil {
					AppendData.Respondent.Name = Value
				}
			}

			// Адрес
			// AppendData.Respondent.Adress, _ = Respondent.InnerText()

			// ИНН
			INN, ErrorINN := Respondent.QuerySelector("div")
			fmt.Println("ErrorINN", ErrorINN)
			if ErrorINN == nil {
				if Value, FindError := INN.InnerText(); FindError != nil {
					Value = strings.ReplaceAll(Value, "ИНН:", "")
					Value = strings.TrimSpace(Value)
					AppendData.Respondent.INN = Value
				}
			}

			fmt.Println(Name.InnerText())
		}

		// Добавляем новую структуру в выходной массив структур
		Datas = append(Datas, AppendData)
	}

	return Datas, nil
}
