package main

import (
	"fmt"
	"net/http"

	"github.com/justinas/nosurf"
)

// WriteToConsole writes something to the console (testing purposes)
func WriteToConsole(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("hit the page")

		// serve the next middleare
		next.ServeHTTP(w, r)
	})
}

// NoSurf adds CSRF production to all POST requests
func NoSurf(next http.Handler) http.Handler {
	csrfHandler := nosurf.New(next)

	csrfHandler.SetBaseCookie(http.Cookie{
		HttpOnly: true,
		Path:     "/",              // cookie path. / is entire site
		Secure:   app.InProduction, // false for dev
		SameSite: http.SameSiteLaxMode,
	})

	return csrfHandler
}

// SessionLoad loads and saves the sessions on every request
func SessionLoad(next http.Handler) http.Handler {
	return session.LoadAndSave(next)
}
