package app

import (
	"fmt"
	"testing"
)

func TestUnwrap(t *testing.T) {
	Input := `1. ООО М4 Б2Б МАРКЕТПЛЕЙС; 0
2. Снегур А. А.; Суд по интеллектуальным правам
3. СИП-344/2023
4. 14.04.2023
5. 14.04.2023
6. c`

	req, reqerr := unwrap(Input)
	if reqerr != nil {
		t.Error(reqerr)
	}
	fmt.Printf("%+v", req)
}
