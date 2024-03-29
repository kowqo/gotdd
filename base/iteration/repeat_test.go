package main

import (
	"fmt"
	"testing"
)

func TestRepeat(t *testing.T) {
	t.Run("Base repeat aaaaa", func(t *testing.T) {
		repeated := Repeat("a", 5)
		expected := "aaaaa"

		if repeated != expected {
			t.Errorf("expected %q but got %q", expected, repeated)
		}
	})

	t.Run("Base repeat bbbbb", func(t *testing.T) {
		repeated := Repeat("b", 5)
		expected := "bbbbb"

		if repeated != expected {
			t.Errorf("expected %q but got %q", expected, repeated)
		}
	})

	t.Run("Base repeat c 10 times", func(t *testing.T) {
		repeated := Repeat("c", 10)
		expected := "cccccccccc"

		if repeated != expected {
			t.Errorf("expected %q but got %q", expected, repeated)
		}
	})

}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 10)
	}
}

func ExampleRepeat() {
	res := Repeat("a", 6)
	fmt.Println(res)
	//Output: aaaaaa
}
