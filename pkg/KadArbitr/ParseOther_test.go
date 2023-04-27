package KadArbitr_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/RB-PRO/KadTG/pkg/KadArbitr"
)

func TestParseCard(t *testing.T) {
	// Создаём ядро
	core, ErrorCore := KadArbitr.NewCore()
	if ErrorCore != nil {
		t.Error(ErrorCore)
	}
	core.Screen("screens/ParseCard1.jpg")

	card, ErrorCard := core.ParseCard("https://kad.arbitr.ru/Card/72197155-c243-47d3-b328-2c421391754a")
	if ErrorCard != nil {
		t.Error(ErrorCard)
	}
	time.Sleep(3 * time.Second)
	core.Screen("screens/ParseCard2.jpg")

	fmt.Printf("%+v", card)

}
