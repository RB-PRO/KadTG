package KadArbitr

import (
	"fmt"
	"testing"
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
