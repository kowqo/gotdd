package main

import (
	"fmt"
	"io"
	"os"
	"time"
)

func main() {
	conf := &ConfSleeper{
		duration: time.Second,
		sleep:    time.Sleep,
	}
	Countdown(os.Stdout, conf)
}

type DefaultSleeper struct {
}

func (d *DefaultSleeper) Sleep() {
	time.Sleep(1 * time.Second)
}

const (
	finalWord      = "Go"
	countdownStart = 3
	sleep          = "sleep"
	write          = "write"
)

func Countdown(out io.Writer, sleeper Sleeper) {
	for i := countdownStart; i > 0; i-- {
		fmt.Fprintln(out, i)
		sleeper.Sleep()
	}
	fmt.Fprintf(out, finalWord)
}

type SpyCountdownOperation struct {
	Calls []string
}

func (s *SpyCountdownOperation) Sleep() {
	s.Calls = append(s.Calls, sleep)
}

func (s *SpyCountdownOperation) Write(p []byte) (n int, e error) {
	s.Calls = append(s.Calls, write)
	return
}

type Sleeper interface {
	Sleep()
}

type SpySleeper struct {
	Calls int
}

func (s *SpySleeper) Sleep() {
	s.Calls++
}

type ConfSleeper struct {
	duration time.Duration
	sleep    func(duration time.Duration)
}

func (c *ConfSleeper) Sleep() {
	c.sleep(c.duration)
}

type SpyTime struct {
	durationSlept time.Duration
}

func (s *SpyTime) Sleep(duration time.Duration) {
	s.durationSlept = duration
}
