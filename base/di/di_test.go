package main

import (
	"bytes"
	"testing"
)

func Test(t *testing.T) {

	assertValue := func(t testing.TB, got, want string) {
		t.Helper()

		if got != want {
			t.Errorf("Want %s, but got %s", want, got)
		}
	}

	t.Run("Базовый тест сравнения значений", func(t *testing.T) {
		buffer := bytes.Buffer{}
		Greet(&buffer, "Chris")

		got := buffer.String()
		want := "Hello, Chris"

		assertValue(t, got, want)
	})

}
