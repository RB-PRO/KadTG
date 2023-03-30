package KadArbitr_test

import (
	"fmt"
	"testing"

	"github.com/RB-PRO/KadArbitr"
)

func TestCouters(t *testing.T) {
	couters, errorReq := KadArbitr.Couters()
	if errorReq != nil {
		t.Error(errorReq)
	}
	fmt.Println(couters)
}

func TestCouters2(t *testing.T) {
	couters, errorReq := KadArbitr.Couters2()
	if errorReq != nil {
		t.Error(errorReq)
	}
	fmt.Println(couters)
}
