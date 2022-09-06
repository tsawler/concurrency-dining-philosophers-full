package main

import (
	"testing"
	"time"
)

func Test_dine(t *testing.T) {
	for i := 0; i < 10; i++ {
		orderFinished = []string{}
		dine()
		if len(orderFinished) != 5 {
			t.Errorf("incorrect length of slice orderFinished; expected 5 but got %d", len(orderFinished))
		}
	}
}

func Test_Main(t *testing.T) {
	for i := 0; i < 10; i++ {
		orderFinished = []string{}
		main()
		if len(orderFinished) != 5 {
			t.Errorf("incorrect length of slice orderFinished; expected 5 but got %d", len(orderFinished))
		}
	}
}

func Test_dineWithVaryingDelays(t *testing.T) {
	var theTests = []struct {
		name  string
		delay time.Duration
	}{
		{"zero delay", time.Second * 0},
		{"half second delay", time.Millisecond * 500},
		{"1 second delay", time.Second},
	}

	for _, e := range theTests {
		// Set orderFinished to an empty slice of strings.
		orderFinished = []string{}

		// Set all sleep times to 0 to speed things along.
		eat = e.delay
		sleepTime = e.delay
		think = e.delay

		// Run the dine function.
		dine()

		// Perform our test.
		if len(orderFinished) != 5 {
			t.Errorf("%s: incorrect length of slice orderFinished; expected 5 but got %d", e.name, len(orderFinished))
		}

		// clean up: set all sleep times to 0 to speed things along.
		eat = 0 * time.Second
		sleepTime = 0 * time.Second
		think = 0 * time.Second
	}
}
