package KadArbitr

import (
	"time"

	"github.com/playwright-community/playwright-go"
)

// Заполнить форму для снятия данных
func (core *CoreReq) FillForm() error {
	// А17-5639/2022
	LocatorFamily, ErrorLocator := core.page.Locator("[placeholder=\"фамилия судьи\"]")
	if ErrorLocator != nil {
		return ErrorLocator
	}

	ClickError := LocatorFamily.Click()
	if ClickError != nil {
		return ClickError
	}

	FillError := LocatorFamily.Fill("А17-5639/2022")
	if FillError != nil {
		return FillError
	}

	return nil
}

func (core *CoreReq) Fill_FilterCases(button string) (ErrorClick error) {

	Selector, ErrorClick := core.page.Locator(`li[class=civil]`)
	if ErrorClick != nil {
		return ErrorClick // could not get entries
	}

	// 7705051215

	ErrorClick = Selector.Click(playwright.PageClickOptions{Button: playwright.MouseButtonLeft})
	if ErrorClick != nil {
		return ErrorClick
	}

	//core.page.Fill(`textarea[placeholder="название, ИНН или ОГРН"]`, "7736050003")

	return nil
}

func (core *CoreReq) FillReqestOne(req Request) (ErrorScreen error) {

	// Если не настроена сущесть стороны, то делаем поиск по "любым" сторонам/лицам
	if req.Part[0].Settings == "" {
		req.Part[0].Settings = "-1"
	}

	// Участник дела, название, ИНН или ОГРН
	if len(req.Part) != 0 {
		core.page.Fill(`div[id=sug-participants] div[class=tag] textarea[placeholder="название, ИНН или ОГРН"]`,
			req.Part[0].Value)
		core.page.SetChecked("div[id=sug-participants] div[class=tag] div#content input[value="+req.Part[0].Settings+"]", true)
	}

	// // Судья, фамилия судьи + инстанция
	// // Пока что не работает. Будет долбавлена позже.
	// core.page.Fill(`div[id=sug-judges] div[class=tag] input[placeholder="фамилия судьи"]`,
	// 	req.Judg[0].Value+", "+req.Judg[0].Instance)
	// fmt.Println(">" + req.Judg[0].Value + ", " + req.Judg[0].Instance + "<")

	// Суд, название суда
	if len(req.Court) != 0 {
		core.page.Fill(`div[id=caseCourt] div[class=tag] input[placeholder="название суда"]`,
			req.Court[0])
	}

	// Номер дела, например, А50-5568/08
	if len(req.Number) != 0 {
		core.page.Fill(`div[id=sug-cases] div[class=tag] input[placeholder="например, А50-5568/08"]`,
			req.Number[0])
	}

	// Дата регистрации дела С
	if len(req.Court) != 0 {
		core.page.Fill(`div[id=sug-dates] label[class=from] input`,
			timeDateToFromFormat(req.DateFrom))
	}

	// Дата регистрации дела ПО
	if len(req.Court) != 0 {
		core.page.Fill(`div[id=sug-dates] label[class=to] input`,
			timeDateToFromFormat(req.DateTo))
	}

	// Судебные поручения
	if req.LegendCheckbox {
		core.page.SetChecked("input[name=WithVKSInstances]", true)
	}

	return nil
}

// Преобразование типа даты от/до
func timeDateToFromFormat(date time.Time) string {
	return date.Format("02.01.2006")
}
