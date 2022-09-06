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
	// set times to 0 to speed things along

	for _, e := range theTests {
		orderFinished = []string{}
		eat = e.delay
		sleepTime = e.delay
		think = 0 * e.delay

		// run the main function
		main()

		if len(orderFinished) != 5 {
			t.Errorf("%s: incorrect length of slice orderFinished; expected 5 but got %d", e.name, len(orderFinished))
		}
	}
}
