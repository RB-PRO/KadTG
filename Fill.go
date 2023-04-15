package KadArbitr

import (
	"io/ioutil"

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

// Выбрать, какой будет запрос. Варианты:
//   - administrative active - Административные
//   - civil - Гражданские
//   - bankruptcy - Банкротсные
//   - "" - Пустой
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

	core.page.Fill(`textarea[placeholder="название, ИНН или ОГРН"]`, "7736050003")

	return nil
}

func (core *CoreReq) Search() (ErrorClick error) {

	core.page.Click("#b-form-submit", playwright.PageClickOptions{
		Delay:   playwright.Float(1),
		Timeout: playwright.Float(3),
		Force:   playwright.Bool(true),
		Strict:  playwright.Bool(true),
	})
	core.page.WaitForSelector(`div[class=b-cases_wrapper]`)

	// save
	html, _ := core.page.QuerySelector("body")
	htmlB, _ := html.InnerHTML()
	err := ioutil.WriteFile("output.html", []byte(htmlB), 0644)
	if err != nil {
		panic(err)
	}

	return nil
}

// Сделать скриншот браузера
func (core *CoreReq) Screen(FileName string) (ErrorScreen error) {
	_, ErrorScreen = core.page.Screenshot(playwright.PageScreenshotOptions{Path: playwright.String(FileName)})
	if ErrorScreen != nil {
		return ErrorScreen
	}
	return nil
}
