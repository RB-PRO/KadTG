package KadArbitr_test

import (
	"testing"

	"github.com/RB-PRO/KadArbitr"
)

func TestNewCore(t *testing.T) {
	core, ErrorCore := KadArbitr.NewCore()
	if ErrorCore != nil {
		t.Error(ErrorCore)
	}
	ErrorStop := core.Stop()
	if ErrorStop != nil {
		t.Error(ErrorStop)
	}
}
