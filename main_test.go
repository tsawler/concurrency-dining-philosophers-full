package main

import (
	"testing"
	"time"
)

func Test_dineWithVaryingDelays(t *testing.T) {
	var theTests = []struct {
		name  string
		delay time.Duration
	}{
		{"zero delay", time.Second * 0},
		{"quarter second delay", time.Millisecond * 250},
		{"half second delay", time.Millisecond * 500},
	}

	for _, e := range theTests {
		// Set orderFinished to an empty slice of strings.
		orderFinished = []string{}

		// Set all sleep times.
		eatTime = e.delay
		sleepTime = e.delay
		thinkTime = e.delay

		// Run the dine function.
		dine()

		// Perform our test.
		if len(orderFinished) != 5 {
			t.Errorf("%s: incorrect length of slice orderFinished; expected 5 but got %d", e.name, len(orderFinished))
		}
	}
}
