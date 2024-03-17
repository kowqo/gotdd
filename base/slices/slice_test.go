package main

import (
	"slices"
	"testing"
)

func TestSum(t *testing.T) {
	t.Run("collection of 5", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5}
		got := Sum(numbers)
		want := 15

		if got != want {
			t.Errorf("want %d but got %d number %v", got, want, numbers)
		}
	})

	t.Run("collection of 4", func(t *testing.T) {
		numbers := make([]int, 4)

		numbers = append(numbers, 1)
		numbers = append(numbers, 10)
		numbers = append(numbers, 100)
		numbers = append(numbers, 1000)

		got := Sum(numbers)
		want := 1111

		if got != want {
			t.Errorf("want %d but got %d number %v", got, want, numbers)
		}
	})

	t.Run("collection of any size 2", func(t *testing.T) {
		numbers := make([]int, 5)
		for i := 0; i < 5; i++ {
			numbers = append(numbers, i)
		}

		got := Sum(numbers)
		want := 10

		if got != want {
			t.Errorf("want %d but got %d number %v", got, want, numbers)
		}
	})
}

func TestSumAll(t *testing.T) {

	t.Run("Sum [1,2] with [4,5]", func(t *testing.T) {
		a := []int{1, 2}
		b := []int{3, 4}
		got := SumAll(a, b)

		want := []int{3, 7}

		if slices.Equal(got, want) {
			t.Errorf("want %d but got %d number 1 - $%v, 2 %v", want, got, a, b)
		}

	})
}

func BenchmarkSum(b *testing.B) {
	for i := 0; i < 100; i++ {
		Sum([]int{123, 123, 123, 123, 123})
	}
}
