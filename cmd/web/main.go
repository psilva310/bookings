package main

import (
	"log"
	"net/http"
	"time"

	"pkg/config"
	"pkg/handlers"
	"pkg/render"

	"github.com/alexedwards/scs/v2"
)

const portNumber = ":8181"

var app config.AppConfig
var session *scs.SessionManager

func main() {

	app.SslEnabled = false

	session = scs.New()
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = app.SslEnabled

	app.Session = session

	// Load pages only once
	templateCache, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("Cannot create template cache")
	}
	app.TemplateCache = templateCache
	app.UseCache = false
	// pass variable address to render.go
	render.NewTemplates(&app)

	// gets the address of the newly created repo
	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)

	// serve
	srv := &http.Server{
		Addr:    portNumber,
		Handler: routes(&app),
	}
	err = srv.ListenAndServe()
	log.Fatal(err)

}
