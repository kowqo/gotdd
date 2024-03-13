package main

import "testing"

func TestHello(t *testing.T) {
	t.Run("if have name Chris, default language", func(t *testing.T) {
		got := Hello("Chris", "")
		want := "Hello " + "Chris"
		assertCorrectMessage(t, got, want)
	})

	t.Run("with no args, default language", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello " + "World!"
		assertCorrectMessage(t, got, want)
	})

	t.Run("args with Russian lang", func(t *testing.T) {
		const name = "Serega"
		got := Hello(name, "RU")
		want := "Privet " + name
		assertCorrectMessage(t, got, want)

	})

	t.Run("args with Russian lang", func(t *testing.T) {
		got := Hello("", "RU")
		want := "Privet " + "Mir"
		assertCorrectMessage(t, got, want)

	})
}

func assertCorrectMessage(t testing.TB, got, want string) {
	t.Helper()

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
