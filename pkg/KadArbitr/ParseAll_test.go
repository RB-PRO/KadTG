package KadArbitr

import (
	"fmt"
	"testing"
	"time"
)

func TestParseAll(t *testing.T) {
	// Создаём ядро
	core, ErrorCore := NewCore()
	if ErrorCore != nil {
		t.Error(ErrorCore)
	}
	core.Screen("screens/ParseAll1.jpg")
	req := Req2() // Создаём запрос на поиск
	// Заполнение формы поиска
	ErrorReq := core.FillReqestOne(req)
	if ErrorReq != nil {
		t.Error(ErrorReq)
	}
	core.Screen("screens/ParseAll2.jpg")

	ErrorSearch := core.Search(req)
	if ErrorSearch != nil {
		t.Error(ErrorSearch)
	}
	core.Screen("screens/ParseAll3.jpg")

	pr, ErrorAll := core.ParseAll()
	if ErrorAll != nil {
		t.Error(ErrorAll)
	}
	fmt.Println(len(pr.Data))

	time.Sleep(2 * time.Second)
	core.Screen("screens/ParseAll4.jpg")
}
