package KadArbitr

import (
	"io/ioutil"
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

// Выбрать, какой будет запрос. Варианты:
//   - administrative active - Административные
//   - civil - Гражданские
//   - bankruptcy - Банкротсные
//   - "" - Пустой
func (core *CoreReq) Fill_FilterCases(button string) (ErrorClick error) {

	Selector, ErrorClick := core.page.Locator(`li[class=civil] i`)
	if ErrorClick != nil {
		return ErrorClick // could not get entries
	}

	ErrorClick = Selector.Click(playwright.PageClickOptions{Button: playwright.MouseButtonLeft})
	if ErrorClick != nil {
		return ErrorClick
	}

	return nil
}

func (core *CoreReq) Search() (ErrorClick error) {
	html, _ := core.page.QuerySelector("body")
	htmlB, _ := html.InnerHTML()
	err := ioutil.WriteFile("output.html", []byte(htmlB), 0644)
	if err != nil {
		panic(err)
	}

	time.Sleep(2 * time.Second)
	core.page.Keyboard().Press("Escape")
	time.Sleep(4 * time.Second)
	core.page.Press("div#b-form-submitters", "Enter", playwright.PagePressOptions{Timeout: playwright.Float(1)})
	time.Sleep(1 * time.Second)
	core.Screen("test2.jpg")
	core.page.Press("div#b-button", "Enter", playwright.PagePressOptions{Timeout: playwright.Float(1)})
	time.Sleep(1 * time.Second)
	core.Screen("test3.jpg")
	core.page.Press("div#b-button-container", "Enter", playwright.PagePressOptions{Timeout: playwright.Float(1)})
	time.Sleep(1 * time.Second)
	core.Screen("test4.jpg")
	core.page.Press("button[alt=Найти]", "Enter", playwright.PagePressOptions{Timeout: playwright.Float(1)})
	time.Sleep(1 * time.Second)
	core.Screen("test5.jpg")
	core.page.Press("div#b-form-submitters span[class=Enter]", "Enter", playwright.PagePressOptions{Timeout: playwright.Float(1)})
	time.Sleep(1 * time.Second)
	core.Screen("test6.jpg")
	core.page.Press("div#b-form-submitters span[class=right]", "Enter", playwright.PagePressOptions{Timeout: playwright.Float(1)})
	time.Sleep(1 * time.Second)

	/*
		selector, ErrorClick := core.page.QuerySelector(`button[alt="Найти"]`)
		if ErrorClick != nil {
			return ErrorClick
		}
		ErrorClick = selector.Click(playwright.ElementHandleClickOptions{Force: playwright.Bool(true)})
		if ErrorClick != nil {
			return ErrorClick
		}
	*/

	// ErrorClick = core.page.DispatchEvent(`input[value="Технические работы"]`, "click", playwright.PageDispatchEventOptions{
	// 	Timeout: playwright.Float(5000),
	// })
	// if ErrorClick != nil {
	// 	return ErrorClick
	// }

	/*
		ErrorClick = core.page.Click("div[class=b-form-submitters] > button[type=submit]",
			playwright.PageClickOptions{Timeout: playwright.Float(5000),
				Force: playwright.Bool(true)})
		if ErrorClick != nil {
			return ErrorClick
		}
	*/

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
