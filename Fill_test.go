package KadArbitr_test

import (
	"fmt"
	"testing"

	"github.com/RB-PRO/KadArbitr"
)

func TestFillForm(t *testing.T) {

	// Создаём ядро
	core, ErrorCore := KadArbitr.NewCore()
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
	core, ErrorCore := KadArbitr.NewCore()
	if ErrorCore != nil {
		t.Error(ErrorCore)
	}

	ErrorScreen := core.Screen("test1.jpg")
	if ErrorScreen != nil {
		t.Error(ErrorScreen)
	}

	// Выбрать административные дела
	ErrorClick := core.Fill_FilterCases("civil")
	if ErrorClick != nil {
		t.Error(ErrorClick)
	}

	ErrorSearch := core.Search()
	if ErrorSearch != nil {
		t.Error(ErrorSearch)
	}

	ErrorScreen2 := core.Screen("test10.jpg")
	if ErrorScreen2 != nil {
		t.Error(ErrorScreen2)
	}

	// Спарсить страницу
	data, ErrorData := core.Parse()
	if ErrorData != nil {
		t.Error(ErrorData)
	}

	fmt.Printf("%+v", data)

	// Останавливаем ядро
	ErrorStop := core.Stop()
	if ErrorStop != nil {
		t.Error(ErrorStop)
	}
}
