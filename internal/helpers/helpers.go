package helpers

import (
	"fmt"
	"net/http"
	"runtime/debug"

	"github.com/cmd-ctrl-q/bookings/internal/config"
)

var app *config.AppConfig

// NewHelpers sets up app config for helpers.
// Gives state-wide access to the app variable.
func NewHelpers(a *config.AppConfig) {
	app = a
}

func ClientError(w http.ResponseWriter, status int) {
	// write to info log
	app.InfoLog.Println("Client error with status of", status)
	http.Error(w, http.StatusText(status), status)
}

func ServerError(w http.ResponseWriter, err error) {
	// debug.Stack() is the stack trace associated with the error message.
	trace := fmt.Sprintf("%s\n%s", err.Error(), debug.Stack())
	app.ErrorLog.Println(trace)

	http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
}

// IsAuthenticated checks if user is authenticated
func IsAuthenticated(r *http.Request) bool {
	exists := app.Session.Exists(r.Context(), "user_id")
	return exists
}
