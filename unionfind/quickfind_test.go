package unionfind

import "testing"

func TestQuickFind(t *testing.T) {
	data := NewQuickFind(10)
	data.Union(5, 0)
	data.Union(6, 5)
	if !data.Connected(5, 0) {
		t.Error("5 and 0 should be connected")
	}
	if !data.Connected(6, 5) {
		t.Error("6 and 5 should be connected")
	}
	if data.Connected(0, 3) {
		t.Error("0 and 3 should not be connected")
	}
}
