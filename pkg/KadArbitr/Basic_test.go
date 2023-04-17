package KadArbitr

import (
	"testing"
)

func TestNewCore(t *testing.T) {
	core, ErrorCore := NewCore()
	if ErrorCore != nil {
		t.Error(ErrorCore)
	}
	ErrorStop := core.Stop()
	if ErrorStop != nil {
		t.Error(ErrorStop)
	}
}
