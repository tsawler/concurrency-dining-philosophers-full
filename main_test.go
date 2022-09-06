package main

import (
	"testing"
	"time"
)

func Test_main(t *testing.T) {
	// set times to 0 to speed things along
	eat = 0 * time.Second
	sleepTime = 0 * time.Second
	think = 0 * time.Second

	// run the main function
	main()

	if len(orderFinished) != 5 {
		t.Error("incorrect length of slice orderFinished")
	}
}
