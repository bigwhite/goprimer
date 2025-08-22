package add

import (
	"testing"
)

func TestAdd(t *testing.T) {
	got := Add(2, 3)
	want := 5
	if got != want {
		t.Errorf("want %d, but got %d", want, got)
	}

	// add zero
	got = Add(2, 0)
	want = 2
	if got != want {
		t.Errorf("want %d, but got %d", want, got)
	}

	// add opposite number
	got = Add(2, -2)
	want = 0
	if got != want {
		t.Errorf("want %d, but got %d", want, got)
	}
}
