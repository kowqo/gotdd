package main

import "testing"

func TestINT(t *testing.T) {
	got := INT()

	want := 2

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
