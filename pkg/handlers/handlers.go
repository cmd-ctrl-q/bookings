package handlers

import (
	"net/http"

	"github.com/cmd-ctrl-q/bookings/pkg/config"
	"github.com/cmd-ctrl-q/bookings/pkg/models"
	"github.com/cmd-ctrl-q/bookings/pkg/render"
)

// Repo the repository used by the handlers
var Repo *Repository

// Repository is the repository type
type Repository struct {
	App *config.AppConfig
}

// NewRepo creates a new repository
func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{
		App: a,
	}
}

// NewHandlers sets the repository for the handlers
func NewHandlers(r *Repository) {
	Repo = r
}

// Home is the handler for the home page
func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {
	// get the remote ip address of the person visiting the site and store in session.
	remoteIP := r.RemoteAddr                              // get ip (version 4 or 6)
	m.App.Session.Put(r.Context(), "remote_ip", remoteIP) // add ip to session
	render.RenderTemplate(w, "home.page.html", &models.TemplateData{})
}

// About is the handler for the about page
func (m *Repository) About(w http.ResponseWriter, r *http.Request) {
	// perform some logic
	stringMap := make(map[string]string)
	stringMap["test"] = "Hello, again"

	// get users ip from session
	remoteIP := m.App.Session.GetString(r.Context(), "remote_ip")

	// store users ip into stringMap
	stringMap["remote_ip"] = remoteIP

	// send data to the template
	render.RenderTemplate(w, "about.page.html", &models.TemplateData{
		StringMap: stringMap,
	})
}
