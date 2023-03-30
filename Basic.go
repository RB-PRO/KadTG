package KadArbitr

import (
	"github.com/playwright-community/playwright-go"
)

const URL string = "https://kad.arbitr.ru/ " // Ссылка на сайт Кад Арбитр

// Базовая стуктура запросов
type CoreReq struct {
	pw      *playwright.Playwright
	browser playwright.Browser
	page    playwright.Page

	Couters map[string]string // Список судов
}

// Создать экземпляр ядра парсинга
func NewCore() (*CoreReq, error) {

	pw, err := playwright.Run()
	if err != nil {
		return nil, err // could not start playwright
	}

	browser, err := pw.Chromium.Launch()
	if err != nil {
		return nil, err // could not launch browser
	}

	page, err := browser.NewPage()
	if err != nil {
		return nil, err // could not create page
	}

	return &CoreReq{
		pw:      pw,
		browser: browser,
		page:    page,
	}, nil
}

// Остановить ядро парсинга
func (core *CoreReq) Stop() error {

	if err := core.browser.Close(); err != nil {
		return err // could not close browser
	}

	if err := core.pw.Stop(); err != nil {
		return err // could not stop Playwright
	}

	return nil
}
