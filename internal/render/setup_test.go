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
