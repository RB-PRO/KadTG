package KadArbitr

import (
	"fmt"
	"testing"
	"time"
)

func TestFillForm(t *testing.T) {

	// Создаём ядро
	core, ErrorCore := NewCore()
	if ErrorCore != nil {
		t.Error(ErrorCore)
	}

	// Останавливаем ядро
	ErrorStop := core.Stop()
	if ErrorStop != nil {
		t.Error(ErrorStop)
	}
}

func TestFill_FilterCases(t *testing.T) {
	// Создаём ядро
	core, ErrorCore := NewCore()
	if ErrorCore != nil {
		t.Error(ErrorCore)
	}
	core.Screen("test1.jpg")

	// Выбрать административные дела
	ErrorClick := core.Fill_FilterCases("civil")
	if ErrorClick != nil {
		t.Error(ErrorClick)
	}

	// ErrorSearch := core.Search()
	// if ErrorSearch != nil {
	// 	t.Error(ErrorSearch)
	// }
	core.Screen("test10.jpg")

	// Спарсить страницу
	data, ErrorData := core.Parse()
	if ErrorData != nil {
		t.Error(ErrorData)
	}

	// fmt.Printf("%+v", data)
	fmt.Printf("Длина массива - %+v", len(data))

	// Останавливаем ядро
	ErrorStop := core.Stop()
	if ErrorStop != nil {
		t.Error(ErrorStop)
	}
}

func TestFiilReq3(t *testing.T) {
	// Создаём ядро
	core, ErrorCore := NewCore()
	if ErrorCore != nil {
		t.Error(ErrorCore)
	}
	core.Screen("screens/Fill_1_Req3.jpg")
	req := Req3() // Создаём запрос на поиск
	// Заполнение формы поиска
	ErrorReq := core.FillReqestOne(req)
	if ErrorReq != nil {
		t.Error(ErrorReq)
	}
	core.Screen("screens/Fill_2_Req3.jpg")

	ErrorSearch := core.Search(req)
	if ErrorSearch != nil {
		t.Error(ErrorSearch)
	}
	core.Screen("screens/Fill_3_Req3.jpg")

	// Останавливаем ядро
	ErrorStop := core.Stop()
	if ErrorStop != nil {
		t.Error(ErrorStop)
	}
}

// Получить тестовый запрос
func Req3() Request {
	return Request{
		// // Стороны
		// Part: []Participant{
		// 	{
		// 		Value:    `7736050003`,   // Истец
		// 		Settings: ParticipantAll, // Категория истца
		// 	},
		// },
		SearchCases: "bankruptcy",
		DateFrom:    time.Date(2023, time.April, 10, 0, 0, 0, 0, time.Local),
		DateTo:      time.Date(2023, time.April, 16, 0, 0, 0, 0, time.Local),
		Court:       []string{"АС города Москвы"},
	}
}
