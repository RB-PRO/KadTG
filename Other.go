package KadArbitr

import "github.com/playwright-community/playwright-go"

// Сделать скриншот браузера
func (core *CoreReq) Screen(FileName string) (ErrorScreen error) {
	_, ErrorScreen = core.page.Screenshot(playwright.PageScreenshotOptions{Path: playwright.String(FileName)})
	if ErrorScreen != nil {
		return ErrorScreen
	}
	return nil
}
