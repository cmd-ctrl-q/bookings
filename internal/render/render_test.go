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

func TestRenderTemplate(t *testing.T) {
	pathToTemplates = "./../../templates"
	// need template cache before we can render templates
	tc, err := CreateTemplateCache()
	if err != nil {
		t.Error(err)
	}

	// if app.UseCache {
	// 	t.Log("app using template cache")
	// }

	// put template cache into app variable
	app.TemplateCache = tc

	// *** render template ***
	// get request
	r, err := getSession()
	if err != nil {
		t.Error(err)
	}

	// get writer
	var w myWriter

	err = Template(&w, r, "home.page.tmpl", &models.TemplateData{})
	if err != nil {
		t.Error("error writing template to browser")
	}

	// check if successfully render a non-existing template
	err = Template(&w, r, "non-existent.page.tmpl", &models.TemplateData{})
	if err == nil {
		t.Error("rendered non-existing template")
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
