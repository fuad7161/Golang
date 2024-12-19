package main

import (
	"encoding/gob"
	"flag"
	"fmt"
	"github.com/alexedwards/scs/v2"
	"github.com/fuad7161/Golang/tree/Project/Bookings/internal/config"
	"github.com/fuad7161/Golang/tree/Project/Bookings/internal/driver"
	"github.com/fuad7161/Golang/tree/Project/Bookings/internal/handlers"
	"github.com/fuad7161/Golang/tree/Project/Bookings/internal/helpers"
	"github.com/fuad7161/Golang/tree/Project/Bookings/internal/models"
	"github.com/fuad7161/Golang/tree/Project/Bookings/internal/render"
	"log"
	"net/http"
	"os"
	"time"
)

var app config.AppConfig
var portNumber = ":8080"
var session *scs.SessionManager
var infoLog *log.Logger
var errorLog *log.Logger

func main() {

	db, err := run()
	if err != nil {
		log.Fatal(err)
	}
	defer db.SQL.Close()
	defer close(app.MailChan)
	listenForMail()

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
	gob.Register(map[string]int{})

	//red flags
	inProduction := flag.Bool("production", true, "Run in production mode")
	useCache := flag.Bool("cache", true, "Use template cache")

	dbHost := flag.String("dbhost", "localhost", "Database host")
	dbName := flag.String("dbname", "postgres", "Database name")
	dbUser := flag.String("dbuser", "postgres", "Database user")
	dbPass := flag.String("dbpass", "", "Database password")
	dbPort := flag.String("dbport", "8080", "Database port")
	dbSSL := flag.String("dbssl", "disable", "Database ssl settings (disable, prefer, require)")

	flag.Parse()

	mailChan := make(chan models.MailData)
	app.MailChan = mailChan

	// Change this to true when in production
	app.InProduction = *inProduction
	app.UseCache = *useCache

	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	app.InfoLog = infoLog

	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)
	app.ErrorLog = errorLog

	session = scs.New()
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = app.InProduction
	app.Session = session

	// connect to database
	log.Println("Connecting to database...")
	connectionString := fmt.Sprintf("host=%s user=%s dbname=%s user=%s password=%s sslmode=%s", *dbHost, *dbUser, *dbName, *dbUser, *dbPass, *dbSSL)
	db, err := driver.ConnectSQL(connectionString)
	if err != nil {
		log.Fatal(err)
	}
	//defer db.SQL.Close()

	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("Cannot create template cache")
		return nil, err

	}
	app.TemplateCache = tc

	repo := handlers.NewRepo(&app, db)
	handlers.NewHandlers(repo)
	render.NewRenderer(&app)
	helpers.NewHelpers(&app)

	fmt.Println(fmt.Sprintf("Starting application on port %s", *dbPort))
	return db, nil
}
