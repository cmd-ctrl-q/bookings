package main

import (
	"encoding/gob"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/alexedwards/scs/v2"
	"github.com/cmd-ctrl-q/bookings/internal/config"
	"github.com/cmd-ctrl-q/bookings/internal/driver"
	"github.com/cmd-ctrl-q/bookings/internal/handlers"
	"github.com/cmd-ctrl-q/bookings/internal/helpers"
	"github.com/cmd-ctrl-q/bookings/internal/models"
	"github.com/cmd-ctrl-q/bookings/internal/render"
)

const portNumber = ":8080"

var app config.AppConfig
var session *scs.SessionManager
var infoLog *log.Logger
var errorLog *log.Logger

// main is the main application function
func main() {
	db, err := run()
	if err != nil {
		log.Fatal(err)
	}
	defer db.SQL.Close()

	defer close(app.MailChan)
	fmt.Println("Starting mail listener...")
	listenForMail()

	fmt.Printf("Starting application on port %s", portNumber)

	srv := &http.Server{
		Addr:    portNumber,
		Handler: routes(&app),
	}

	err = srv.ListenAndServe()
	log.Fatal(err)
}

func run() (*driver.DB, error) {
	// what am I going to put in the session
	gob.Register(models.Reservation{})
	gob.Register(models.User{})
	gob.Register(models.Room{})
	gob.Register(models.Restriction{})

	mailChan := make(chan models.MailData)
	app.MailChan = mailChan

	// change this to true when in production
	app.InProduction = false

	infoLog = log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	app.InfoLog = infoLog

	errorLog = log.New(os.Stdout, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)
	app.ErrorLog = errorLog

	session = scs.New()
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = app.InProduction

	app.Session = session

	// connect to database
	log.Println("Connecting to database...")
	db, err := driver.ConnectSQL("host=localhost port=5432 dbname=bookings user=plutonium password= sslmode=disable")
	if err != nil {
		log.Fatal("Cannot connect to database! Dying...")
	}
	log.Println("Connected to database!")

	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("cannot create template cache")
		return nil, err
	}

	app.TemplateCache = tc
	app.UseCache = false

	repo := handlers.NewRepo(&app, db)
	handlers.NewHandlers(repo)
	render.NewRenderer(&app)
	helpers.NewHelpers(&app)

	return db, nil
}

// const (
// 	portNumber = ":8080"
// )

// var (
// 	app     config.AppConfig
// 	session *scs.SessionManager
// 	infoLog *log.Logger
// 	errLog  *log.Logger
// )

// func main() {

// 	db, err := run()
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	defer db.SQL.Close()

// 	// what am I going to put in the session
// 	gob.Register(models.Reservation{})

// 	// initialize loggers
// 	// info logger
// 	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
// 	app.InfoLog = infoLog
// 	// error logger
// 	errLog = log.New(os.Stdout, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)
// 	app.ErrorLog = errLog

// 	// initialize sessions
// 	session = scs.New()
// 	session.Lifetime = 24 * time.Hour // let sessions last for 24 hours
// 	// should the cookie persist after the client browser window is closed by the end user?
// 	session.Cookie.Persist = true
// 	session.Cookie.SameSite = http.SameSiteLaxMode
// 	session.Cookie.Secure = app.InProduction // encrypted cookie. false for development

// 	app.Session = session

// 	tc, err := render.CreateTemplateCache()
// 	if err != nil {
// 		log.Fatal("cannot create template cache")
// 	}

// 	app.TemplateCache = tc
// 	app.UseCache = false

// 	repo := handlers.NewRepo(&app, db)
// 	handlers.NewHandlers(repo)

// 	render.NewRenderer(&app)

// 	fmt.Printf("Staring application on port %s\n", portNumber)

// 	// router
// 	router := routes(&app)

// 	srv := &http.Server{
// 		Addr:    portNumber,
// 		Handler: router,
// 	}

// 	err = srv.ListenAndServe()
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// }

// func run() (*driver.DB, error) {
// 	// what am I going to put in the session
// 	gob.Register(models.Reservation{})
// 	gob.Register(models.User{})
// 	gob.Register(models.Room{})
// 	gob.Register(models.RoomRestriction{})

// 	// change this to true when in production (there is a better way to do this)
// 	app.InProduction = false

// 	// initialize sessions
// 	session = scs.New()
// 	session.Lifetime = 24 * time.Hour // let sessions last for 24 hours
// 	// should the cookie persist after the client browser window is closed by the end user?
// 	session.Cookie.Persist = true
// 	session.Cookie.SameSite = http.SameSiteLaxMode
// 	session.Cookie.Secure = app.InProduction // encrypted cookie. false for development

// 	app.Session = session

// 	// connect to db
// 	log.Println("Connecting to db...")
// 	db, err := driver.ConnectSQL("host=localhost port=5432 dbname=bookings user=plutonium password= sslmode=disable")
// 	if err != nil {
// 		log.Fatal("Cannot connect to database! Dying...")
// 	}
// 	// defer db.SQL.Close()
// 	log.Println("Connected to database!")

// 	tc, err := render.CreateTemplateCache()
// 	if err != nil {
// 		log.Fatal("cannot create template cache")
// 		return nil, err
// 	}

// 	app.TemplateCache = tc
// 	app.UseCache = false

// 	repo := handlers.NewRepo(&app, db)
// 	handlers.NewHandlers(repo)
// 	render.NewRenderer(&app)
// 	helpers.NewHelpers(&app)

// 	return db, nil
// }
