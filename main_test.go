package main

import (
	"testing"
	"time"
)

func Test_main(t *testing.T) {
	eat = 0 * time.Second
	sleepTime = 0 * time.Second
	think = 0 * time.Second

	main()

	if len(orderFinished) != 5 {
		t.Error("incorrect length of slice orderFinished")
	}
	orderFinished = []string{}
}
