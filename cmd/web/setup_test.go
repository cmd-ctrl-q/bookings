package main

import (
	"net/http"
	"os"
	"testing"
)

// runs before tests run

func TestMain(m *testing.M) {

	// run the tests and exit
	os.Exit(m.Run())
}

type myHandler struct{}

func (mh *myHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {}
