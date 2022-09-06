package main

import (
	"os"
	"testing"
	"time"
)

func TestMain(m *testing.M) {
	// Set all sleep times to 0 to speed things along.
	eat = 0 * time.Second
	sleepTime = 0 * time.Second
	think = 0 * time.Second

	os.Exit(m.Run())
}
