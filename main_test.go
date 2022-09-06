package main

import (
	"testing"
	"time"
)

func Test_main(t *testing.T) {
	var theTests = []struct {
		name  string
		delay time.Duration
	}{
		{"zero delay", time.Second * 0},
		{"half second delay", time.Millisecond * 500},
		{"1 second delay", time.Second},
	}

	for _, e := range theTests {
		// set orderFinished to an empty slice of strings
		orderFinished = []string{}

		// set times to 0 to speed things along
		eat = e.delay
		sleepTime = e.delay
		think = 0 * e.delay

		// run the dine function
		dine()

		// perform our tests
		if len(orderFinished) != 5 {
			t.Errorf("%s: incorrect length of slice orderFinished; expected 5 but got %d", e.name, len(orderFinished))
		}
	}
}

func Test_dine(t *testing.T) {
	for k := 0; k < 10; k++ {
		orderFinished = []string{}
		dine()
		if len(orderFinished) != 5 {
			t.Errorf("incorrect length of slice orderFinished; expected 5 but got %d", len(orderFinished))
		}
	}
}
