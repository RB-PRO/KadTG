package KadArbitr

import (
	"strings"
	"time"

	"github.com/playwright-community/playwright-go"
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

	// Содержимое дела(Если провалить в UrlNumber)
	// см. файл ParseOther.go
	Card Card
}

// Сторона конфликта
type Side struct {
	Name   string // Название компании
	Adress string // Адресс
	INN    string // ИНН
}

func (core *CoreReq) Parse() ([]Data, error) {
	Datas := make([]Data, 0)

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
		// * Дата *
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

		// * Номер дела + ссылка на дело *
		NumberSelector, NumberError := TD[0].QuerySelector("a")
		if NumberError == nil {
			// * Номер дела *
			NumberText, FindNumberText := NumberSelector.InnerText()
			if FindNumberText == nil {
				AppendData.Number = NumberText
			}

			// * Ссылка на дело *
			HrefStr, IsHref := NumberSelector.GetAttribute("href")
			if IsHref == nil {
				AppendData.UrlNumber = HrefStr
			}
		}

		// ************* //
		// * 2 колонка * //
		// ************* //
		// * Судья *
		Judge, ErrorJudge := TD[1].QuerySelector("div div.judge")
		if ErrorJudge == nil {
			JudgeText, IsJudge := Judge.InnerText()
			if IsJudge == nil {
				AppendData.Judge = JudgeText
			}
		}

		// * Инстанция *
		Court, ErrorCourt := TD[1].QuerySelector("div div:last-of-type")
		if ErrorCourt == nil {
			Instance, ErrorInnerCourt := Court.InnerText()
			if ErrorInnerCourt != nil {
				AppendData.Instance = Instance
			}
		}

		// ************* //
		// * 3 колонка * //
		// ************* //
		// * Истец *
		Plaintiff, ErrorPlaintiff := TD[2].QuerySelector("span[class=js-rolloverHtml]")
		if ErrorPlaintiff == nil {
			AppendData.Plaintiff = plaintiff_2_respondent(Plaintiff)
		}

		// ************* //
		// * 4 колонка * //
		// ************* //
		// * Ответчик *
		Respondent, ErrorRespondent := TD[3].QuerySelector("span[class=js-rolloverHtml]")
		if ErrorRespondent == nil {
			AppendData.Respondent = plaintiff_2_respondent(Respondent)
		}

		// Добавляем новую структуру в выходной массив структур
		Datas = append(Datas, AppendData)
	}

	return Datas, nil
}

// Распарсить структуру для Истца и ответчика
// Распарсивает клетку в структуру Side, которую и возвращает
func plaintiff_2_respondent(side playwright.ElementHandle) (OutPutSide Side) {
	// * Название компании *
	Name, ErrorName := side.QuerySelector("strong")
	if ErrorName == nil {
		Name, ErrorInnerName := Name.InnerText()
		if ErrorInnerName == nil {
			OutPutSide.Name = Name
		}
	}

	// * Адрес *
	Adress, ErrorInnerAdress := side.InnerText()
	if ErrorInnerAdress == nil {
		OutPutSide.Adress = Adress
	}

	// * ИНН *
	INN, _ := side.QuerySelector("div")
	//fmt.Println("INN", INN)
	if INN != nil {
		INN, ErrorInnerINN := INN.InnerText()
		if ErrorInnerINN == nil {
			INN = strings.ReplaceAll(INN, "ИНН:", "")
			OutPutSide.INN = strings.TrimSpace(INN)
		}
	}
	return OutPutSide
}
