package KadArbitr_test

import (
	"testing"

	"github.com/RB-PRO/KadArbitr"
)

func TestCouters(t *testing.T) {

	// Создаём ядро
	core, ErrorCore := KadArbitr.NewCore()
	if ErrorCore != nil {
		t.Error(ErrorCore)
	}

	// Тестрируем парсинг списка судов
	errorReq := core.ParseCouters()
	if errorReq != nil {
		t.Error(errorReq)
	}
	if len(core.Couters) == 0 {
		t.Error("Не смогу собрать список судов. Найдено судов - нуль.")
	}

	// Останавливаем ядро
	ErrorStop := core.Stop()
	if ErrorStop != nil {
		t.Error(ErrorStop)
	}
}
