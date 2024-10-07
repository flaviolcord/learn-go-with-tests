package main

import (
	"fmt"
	"io"
	"os"
	"time"
)

const (
  finalWord = "Go!"
  countDownStart = 3
)

type Sleeper interface {
  Sleep()
}

type ConfigurableSleeper struct {
  duration time.Duration
  sleep func(time.Duration)
}

func (cs ConfigurableSleeper) Sleep() {
  cs.sleep(cs.duration)
}

func CountDown(w io.Writer, sleeper Sleeper) {
  for i := countDownStart; i > 0; i-- {
    fmt.Fprintln(w, i)
    sleeper.Sleep()
  }
  fmt.Fprint(w, finalWord)
}

func main() {
  sleeper := ConfigurableSleeper{1 * time.Second, time.Sleep}
  CountDown(os.Stdout, sleeper)
}
