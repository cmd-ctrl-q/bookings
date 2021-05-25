package handlers

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
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
	// get requests
	{"home", "/", "GET", []postData{}, http.StatusOK},
	{"about", "/about", "GET", []postData{}, http.StatusOK},
	{"gq", "/generals-quarters", "GET", []postData{}, http.StatusOK},
	{"ms", "/majors-suite", "GET", []postData{}, http.StatusOK},
	{"sa", "/search-availability", "GET", []postData{}, http.StatusOK},
	{"contact", "/contact", "GET", []postData{}, http.StatusOK},
	{"mr", "/make-reservation", "GET", []postData{}, http.StatusOK},

	// post requests
	{"post-search-avail", "/search-availability", "POST", []postData{
		{key: "start", value: "2021-01-01"},
		{key: "start", value: "2021-01-02"},
	}, http.StatusOK},
	{"post-search-avail-json", "/search-availability-json", "POST", []postData{
		{key: "start", value: "2021-01-01"},
		{key: "start", value: "2021-01-02"},
	}, http.StatusOK},
	{"post-make-reservation", "/make-reservation", "POST", []postData{
		{key: "first_name", value: "John"},
		{key: "last_name", value: "Wick"},
		{key: "email", value: "john.wick@jw.com"},
		{key: "phone", value: "555-555-5555"},
	}, http.StatusOK},
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
				t.Errorf("for %s, expected %d but got %d", e.name, e.expectedStatusCode, resp.StatusCode)
			}
		}
	}
}
