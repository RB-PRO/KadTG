package KadArbitr

import (
	"fmt"
	"testing"
)

func TestParse(t *testing.T) {
	// Создаём ядро
	core, ErrorCore := NewCore()
	if ErrorCore != nil {
		t.Error(ErrorCore)
	}

	// Заполнить ввод данных и выполнить запрос
	FillError := core.FillForm()
	if FillError != nil {
		t.Error(FillError)
	}

	// Спарсить страницу
	data, ErrorData := core.Parse()
	if ErrorData != nil {
		t.Error(ErrorData)
	}

	fmt.Println(data)

	// Останавливаем ядро
	ErrorStop := core.Stop()
	if ErrorStop != nil {
		t.Error(ErrorStop)
	}
}
