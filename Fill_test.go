package KadArbitr_test

import (
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
