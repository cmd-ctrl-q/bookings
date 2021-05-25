package main

import "testing"

// TestRun tests the run() in main
func TestRun(t *testing.T) {
	err := run()
	if err != nil {
		t.Error("failed run()")
	}
}
