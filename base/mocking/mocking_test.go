package main

import (
	"bytes"
	"reflect"
	"testing"
	"time"
)

func TestCountdown(t *testing.T) {

	assertValue := func(t testing.TB, got, want string) {
		t.Helper()

		if got != want {
			t.Errorf("want '%s' but got '%s'", want, got)
		}
	}

	assertSpyCallsCount := func(t testing.TB, got, want int) {
		t.Helper()

		if got != want {
			t.Errorf("want %d calls but got %d calls", want, got)
		}
	}
	t.Run("Sleep before every write", func(t *testing.T) {
		spySleepPrinter := &SpyCountdownOperation{}

		Countdown(spySleepPrinter, spySleepPrinter)

		want := []string{
			write,
			sleep,
			write,
			sleep,
			write,
			sleep,
			write,
		}

		if !reflect.DeepEqual(want, spySleepPrinter.Calls) {
			t.Errorf("wanted calls %v got %v", want, spySleepPrinter.Calls)
		}
	})

	t.Run("Base countdown", func(t *testing.T) {

		buffer := &bytes.Buffer{}
		spySleeper := &SpySleeper{}

		Countdown(buffer, spySleeper)

		got := buffer.String()
		want := `3
2
1
Go`

		assertValue(t, got, want)

		assertSpyCallsCount(t, spySleeper.Calls, 3)
	})

	t.Run("test conf sleeper", func(t *testing.T) {
		sleepTime := time.Second * 5

		spyTime := &SpyTime{}

		sleeper := ConfSleeper{sleepTime, spyTime.Sleep}

		sleeper.Sleep()

		if spyTime.durationSlept != sleepTime {
			t.Errorf("should have slept for %v but slept for %v", sleepTime, spyTime.durationSlept)
		}
	})
}
