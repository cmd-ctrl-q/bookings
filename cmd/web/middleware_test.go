package main

import (
	"fmt"
	"net/http"
	"testing"
)

// set up env before running test

func TestNoSurf(t *testing.T) {
	var myH myHandler

	h := NoSurf(&myH)

	// test if it returns a handler
	// store in v, whatever type h is
	switch v := h.(type) {
	case http.Handler:
		// do nothing (expected)
	default:
		t.Error(fmt.Sprintf("wants: http.Handler, has: %T", v))
	}
}

func TestSessionLoad(t *testing.T) {
	var myH myHandler

	h := SessionLoad(&myH)

	switch v := h.(type) {
	case http.Handler:
		// do nothing (expected)
	default:
		t.Error(fmt.Sprintf("wants: http.Handler, has: %T", v))
	}
}
