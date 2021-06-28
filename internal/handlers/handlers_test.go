package handlers

import (
	"context"
	"log"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/cmd-ctrl-q/bookings/internal/models"
)

type postData struct {
	key   string
	value string
}

// var for test
var theTests = []struct {
	name   string
	url    string
	method string
	params []postData

	// to check if test has passed
	expectedStatusCode int
}{
	// // get requests
	// {"home", "/", "GET", []postData{}, http.StatusOK},
	// {"about", "/about", "GET", []postData{}, http.StatusOK},
	// {"gq", "/generals-quarters", "GET", []postData{}, http.StatusOK},
	// {"ms", "/majors-suite", "GET", []postData{}, http.StatusOK},
	// {"sa", "/search-availability", "GET", []postData{}, http.StatusOK},
	// {"contact", "/contact", "GET", []postData{}, http.StatusOK},
	// {"mr", "/make-reservation", "GET", []postData{}, http.StatusOK},

	// // post requests
	// {"post-search-avail", "/search-availability", "POST", []postData{
	// 	{key: "start", value: "2021-01-01"},
	// 	{key: "start", value: "2021-01-02"},
	// }, http.StatusOK},
	// {"post-search-avail-json", "/search-availability-json", "POST", []postData{
	// 	{key: "start", value: "2021-01-01"},
	// 	{key: "start", value: "2021-01-02"},
	// }, http.StatusOK},
	// {"post-make-reservation", "/make-reservation", "POST", []postData{
	// 	{key: "first_name", value: "John"},
	// 	{key: "last_name", value: "Wick"},
	// 	{key: "email", value: "john.wick@jw.com"},
	// 	{key: "phone", value: "555-555-5555"},
	// }, http.StatusOK},
}

// table test
func TestHandlers(t *testing.T) {

	// get routes
	routes := getRoutes()

	// create a test webserver that returns a status code
	ts := httptest.NewTLSServer(routes)
	defer ts.Close()

	// create client to call the server.
	for _, e := range theTests {
		// get request tests
		switch e.method {
		case "GET":
			resp, err := ts.Client().Get(ts.URL + e.url)
			if err != nil {
				t.Log(err)
				t.Fatal(err)
			}

			if resp.StatusCode != e.expectedStatusCode {
				t.Errorf("for %s, expected %d but got %d", e.name, e.expectedStatusCode, resp.StatusCode)
			}
			continue
		case "POST":
			// create url.Values{} object
			values := url.Values{}
			//  populate url.Values{} with the POST parameters
			for _, x := range e.params {
				values.Add(x.key, x.value)
			}

			resp, err := ts.Client().PostForm(ts.URL+e.url, values)
			if err != nil {
				t.Log(err)
				t.Fatal(err)
			}

			if resp.StatusCode != e.expectedStatusCode {
				// t.Errorf("for %s, expected %d but got %d", e.name, e.expectedStatusCode, resp.StatusCode)
				t.Errorf("for #{e.name}, expected #{e.expectedStatusCode} but got #{resp.StatusCode}")
			}
		}
	}
}

func TestRepository_Reservation(t *testing.T) {
	reservation := models.Reservation{
		RoomID: 1,
		Room: models.Room{
			ID:       1,
			RoomName: "General's Quarters",
		},
	}

	// create a test request with body = nil
	req, _ := http.NewRequest("GET", "/main-reservation", nil)
	// get context
	ctx := getCtx(req)
	// add context to our test request
	req = req.WithContext(ctx)

	// Create request recorder
	// A new recorder simulates an entire rest system.
	// The request from the users browser to the response from the backend.
	// User passes handler a request, gets a response writer,
	// then response writer writes the response to the web browser
	rr := httptest.NewRecorder()
	// put reservation into the session
	session.Put(ctx, "reservation", reservation)
	// Convert handler reservation to a handler function.
	// To be able to call directly.
	handler := http.HandlerFunc(Repo.Reservation)
	// call reservation function
	handler.ServeHTTP(rr, req)

	// check if test passes
	if rr.Code == http.StatusOK {
		t.Errorf("Reservation handler returned wrong response code: got %d, wanted %d", rr.Code, http.StatusOK)
	}
}

func getCtx(req *http.Request) context.Context {
	ctx, err := session.Load(req.Context(), req.Header.Get("X-Session"))
	if err != nil {
		log.Println(err)
	}

	return ctx
}
