package models

import (
	"github.com/cmd-ctrl-q/bookings/internal/forms"
)

type TemplateData struct {
	StringMap map[string]string      `json:"string_map"`
	IntMap    map[string]int         `json:"int_map"`
	FloatMap  map[string]float32     `json:"float_map"`
	Data      map[string]interface{} `json:"data"`

	// Cross-site Request Forgery Token
	CSRFToken string `json:"csrf_token"`

	// Flash, Warning, and Error are used to send messages to the user
	Flash   string `json:"flash"`
	Warning string `json:"warning"`
	Error   string `json:"error"`

	// Form is the html form
	Form *forms.Form `json:"form"`

	// if value > 0 then user is logged in.
	// if value < 0 then user is not logged in.
	IsAuthenticated int
}
