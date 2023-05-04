package KadArbitr

import (
	"errors"
	"math"

	"github.com/playwright-community/playwright-go"
)

// Из общего к-ва записей найти к-во страниц
// На одной странице максимум 25 записей
// к-во страниц не может быть больше 40,
// соответственноесли записей больше или равно 1000, то
func NumberTotalPages(count int) (pages int) {

	// Если количество записей выходит за допустимый диапазон
	if count >= 1000 {
		return 40
	}

	// Делаем преобразование
	pages = int(math.Ceil(float64(count) / 25.0))

	return pages
}

// Ошибка, которая информирует о том, что нет кнопке далее или вы упёрлись в последнюю страницу
var ErrorNotIsVisible error = errors.New("NextPage: Не вижу кнопку 'Ctrl ->', которая открывает следующую страницу")

// Выбрать следующую страницу и подождать, пока она загрузится
//
// Нажимаем на кнопку "далее"
func (core *CoreReq) NextPage() (ErrorClick error) {

	var IsVisible bool // переменная, которая имеет сведения о том, последняя ли это страница

	// Если кнопка на следующую ссылку активна
	if IsVisible, ErrorClick = core.page.IsVisible("ul[id=pages] li[class=rarr] a"); ErrorClick != nil {
		return ErrorClick
	}

	// Если кнопка видна
	if !IsVisible {
		return ErrorNotIsVisible
	}

	// Нажимаем на кнопку "далее"
	if ErrorClick = core.page.Click("ul[id=pages] li[class=rarr]", playwright.PageClickOptions{
		Force:   playwright.Bool(true),
		Delay:   playwright.Float(100),
		Timeout: playwright.Float(5000),
	}); ErrorClick != nil {
		return ErrorClick
	}

	return nil
}
