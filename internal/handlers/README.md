

```go
// Home is the handler for the home page
func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {
	// get the remote ip address of the person visiting the site and store in session.
	remoteIP := r.RemoteAddr                              // get ip (version 4 or 6)
	m.App.Session.Put(r.Context(), "remote_ip", remoteIP) // add ip to session
	render.RenderTemplate(w, r, "home.page.tmpl", &models.TemplateData{})
}
```