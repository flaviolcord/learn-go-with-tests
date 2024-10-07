package main

import (
	"bytes"
	"reflect"
	"testing"
	"time"
)

const (
  sleep = "sleep"
  write = "write"
)

func TestCountDown(t *testing.T) {
  t.Run("verify the number of sleep calls", func(t *testing.T) {
    buffer := &bytes.Buffer{}
    spySleeper := &SpySleeper{}

    CountDown(buffer, spySleeper)

    got := buffer.String()
    want := "3\n2\n1\nGo!"

    if got != want {
      t.Errorf("got: %s, want: %s", got, want)
    }

    if spySleeper.Calls != 3 {
      t.Errorf("not enough calls to sleeper, want 3 and got: %d", spySleeper.Calls)
    }
  })

  t.Run("sleep before every print", func(t *testing.T) {
    spySpleepPrinter := &SpyCountDownOperations{}
    CountDown(spySpleepPrinter, spySpleepPrinter)

    want := []string{
      write,
      sleep,
      write,
      sleep,
      write,
      sleep,
      write,
    } 

    if !reflect.DeepEqual(want, spySpleepPrinter.Calls) {
      t.Errorf("wanted calls %v, and got: %v", want, spySpleepPrinter.Calls)
    }
  })
}

func TestConfigurableSleeper(t *testing.T) {
  sleepTime := 5 * time.Second
  spyTime := &SpyTime{}

  sleeper := ConfigurableSleeper{sleepTime, spyTime.Sleep}
  sleeper.Sleep()

  if spyTime.durationSlept != sleepTime {
    t.Errorf("should be slept for: %v but slept for: %v", sleepTime, spyTime.durationSlept)
  }
}

type SpyCountDownOperations struct {
  Calls []string
}

func (s *SpyCountDownOperations) Sleep() {
  s.Calls = append(s.Calls, sleep)
}

func (s *SpyCountDownOperations) Write(p []byte) (n int, err error) {
  s.Calls = append(s.Calls, write)
  return
}

type SpyTime struct {
  durationSlept time.Duration
}

func (st *SpyTime) Sleep(duration time.Duration) {
  st.durationSlept = duration
}

type SpySleeper struct {
  Calls int
}

func (s *SpySleeper) Sleep() {
  s.Calls++
}
