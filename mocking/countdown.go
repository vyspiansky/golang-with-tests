package main

import (
	"fmt"
	"io"
	"iter"
	"os"
	"time"
)

// func Countdown(out *bytes.Buffer) {
// 	fmt.Fprint(out, "3")
// }

const finalWord = "Go!"

// const countdownStart = 3

type Sleeper interface {
	Sleep()
}

type ConfigurableSleeper struct {
	duration time.Duration
	sleep    func(time.Duration)
}

func (c *ConfigurableSleeper) Sleep() {
	c.sleep(c.duration)
}

// A simple iterator, which will yield numbers in reverse order.
// iter.Seq[T] type is a type alias for func(func(T) bool)
func countDownFrom(from int) iter.Seq[int] {
	return func(yield func(int) bool) {
		for i := from; i > 0; i-- {
			if !yield(i) {
				return
			}
		}
	}
}

// Use a general purpose interface instead
func Countdown(out io.Writer, sleeper Sleeper) {
	// for i := countdownStart; i > 0; i-- {
	for i := range countDownFrom(3) {
		fmt.Fprintln(out, i)

		// 1-second pause
		// time.Sleep(1 * time.Second)
		sleeper.Sleep()
	}
	fmt.Fprint(out, finalWord)
}

// type DefaultSleeper struct{}

// func (d *DefaultSleeper) Sleep() {
// 	time.Sleep(1 * time.Second)
// }

func main() {
	sleeper := &ConfigurableSleeper{1 * time.Second, time.Sleep}
	Countdown(os.Stdout, sleeper)
}
