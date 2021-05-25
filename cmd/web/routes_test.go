package main

import (
	"fmt"
	"testing"

	"github.com/cmd-ctrl-q/bookings/internal/config"
	"github.com/go-chi/chi"
)

func TestRoutes(t *testing.T) {
	var app config.AppConfig

	mux := routes(&app)

	switch v := mux.(type) {
	case *chi.Mux:
		// do nothing
	default:
		t.Error(fmt.Sprintf("wants: http.Handler has: %T", v))
	}
}
