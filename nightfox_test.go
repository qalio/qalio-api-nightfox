package nightfox

import "testing"

func TestInit(t *testing.T) {
	expected := "nightfox module initialized"
	if got := Init(); got != expected {
		t.Errorf("Init() = %q, want %q", got, expected)
	}
}
