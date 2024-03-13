package main

import "testing"

func TestHello(t *testing.T) {
	t.Run("if have name Chris, EN language", func(t *testing.T) {
		got := Hello("Chris", "EN")
		want := "Hello " + "Chris"
		assertCorrectMessage(t, got, want)
	})

	t.Run("with no args, EN language", func(t *testing.T) {
		got := Hello("", "EN")
		want := "Hello " + "World!"
		assertCorrectMessage(t, got, want)
	})

	t.Run("args with Russian lang", func(t *testing.T) {
		const name = "Serega"
		got := Hello(name, "RU")
		want := "Privet " + name
		assertCorrectMessage(t, got, want)

	})

	t.Run("no args with Russian lang", func(t *testing.T) {
		got := Hello("", "RU")
		want := "Privet " + "World!"
		assertCorrectMessage(t, got, want)

	})

	t.Run("args with French lang", func(t *testing.T) {
		got := Hello("Serega", "FR")
		want := "Bonjour " + "Serega"
		assertCorrectMessage(t, got, want)

	})

	t.Run("no args with French lang", func(t *testing.T) {
		got := Hello("", "FR")
		want := "Bonjour " + "World!"
		assertCorrectMessage(t, got, want)

	})
}

func assertCorrectMessage(t testing.TB, got, want string) {
	t.Helper()

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
