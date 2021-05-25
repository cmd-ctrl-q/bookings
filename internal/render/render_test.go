package render

import (
	"net/http"
	"testing"

	"github.com/cmd-ctrl-q/bookings/internal/models"
)

func TestAddDefaultData(t *testing.T) {
	var td models.TemplateData
	// add context info to session
	r, err := getSession()
	if err != nil {
		t.Error(err)
	}

	// add data to session
	session.Put(r.Context(), "flash", "123")

	// call the function with the session info in it
	result := AddDefaultData(&td, r)
	if result.Flash != "123" {
		t.Error("flash value of 123 not found in session")
	}
}

func getSession() (*http.Request, error) {
	// make a request to some url
	r, err := http.NewRequest("GET", "/some-url", nil)
	if err != nil {
		return nil, err
	}

	ctx := r.Context()
	// add X-Session to context
	ctx, _ = session.Load(ctx, r.Header.Get("X-Session"))
	r = r.WithContext(ctx)

	return r, nil
}
