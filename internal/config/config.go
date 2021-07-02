package config

import (
	"html/template"
	"log"

	"github.com/alexedwards/scs/v2"
	"github.com/cmd-ctrl-q/bookings/internal/models"
)

// AppConfig holds the application config
type AppConfig struct {
	UseCache bool

	// A map to cache templates
	TemplateCache map[string]*template.Template

	// Built-in loggers
	InfoLog  *log.Logger
	ErrorLog *log.Logger

	// Set to true if in production
	InProduction bool
	Session      *scs.SessionManager

	// MailChan is is the mailing channel
	MailChan chan models.MailData
}
