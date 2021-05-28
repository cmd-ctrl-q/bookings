package render

import (
	"encoding/gob"
	"net/http"
	"os"
	"testing"
	"time"

	"github.com/alexedwards/scs/v2"
	"github.com/cmd-ctrl-q/bookings/internal/config"
	"github.com/cmd-ctrl-q/bookings/internal/models"
)

var session *scs.SessionManager
var testApp config.AppConfig

// testing.M is called before any tests are run
func TestMain(m *testing.M) {
	// what am I going to put in the session
	gob.Register(models.Reservation{})

	// change this to true when in production (there is a better way to do this)
	testApp.InProduction = false

	session = scs.New()
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = false

	testApp.Session = session

	// easy way to mock app, ie make new object testApp and set it as a pointer to app
	app = &testApp

	os.Exit(m.Run())
}

type myWriter struct{}

func (w *myWriter) Header() http.Header {
	return http.Header{}
}

func (w *myWriter) Write(b []byte) (int, error) {
	return len(b), nil
}

func (w *myWriter) WriteHeader(statusCode int) {}

func TestNewTemplates(t *testing.T) {
	NewTemplates(app)
}

func TestCreateTemplateCache(t *testing.T) {
	_, err := CreateTemplateCache()
	if err != nil {
		t.Error(err)
	}

}
