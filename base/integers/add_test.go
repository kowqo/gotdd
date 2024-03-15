package main

import (
	"fmt"
	"testing"
)

func ExampleAdd() {
	sum := Add(1, 5)
	fmt.Println(sum)
	// Output 6
}

func TestAdd(t *testing.T) {
	t.Run("2 plus 5", func(t *testing.T) {
		got := Add(2, 5)
		want := 7
		assertCorrect(t, got, want)
	})

	t.Run("0 plus 10", func(t *testing.T) {
		got := Add(0, 10)
		want := 10
		assertCorrect(t, got, want)
	})

}

func assertCorrect(t testing.TB, got, want int) {
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
