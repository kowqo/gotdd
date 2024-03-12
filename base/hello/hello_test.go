package main

import "testing"

func TestHello(t *testing.T) {
	got := Hello("Chris")

	want := "Hello"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
