package main

import (
	"github.com/alexedwards/scs/v2"
	"github.com/cozmus/hello-world/pkg/config"
	"github.com/cozmus/hello-world/pkg/handlers"
	"github.com/cozmus/hello-world/pkg/render"
	"log"
	"net/http"
	"time"
)

const portNumber = ":8080"
var app config.AppConfig
var session *scs.SessionManager

func main() {
	// change to true in prod
	app.InProduction = false // true for prod

	session = scs.New()
	session.Lifetime = 24 * time.Hour

	// if browser closed session will persist
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = app.InProduction

	app.Session = session

	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("cannot create template cache")
	}

	app.TemplateCache = tc
	app.UseCache = false

	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)

	render.NewTemplates(&app)

	//http.HandleFunc("/", handlers.Repo.Home)
	//http.HandleFunc("/about", handlers.Repo.About)


	log.Println("Starting app on port:", portNumber)
	//_ = http.ListenAndServe(portNumber, nil)

	srv := &http.Server{
		Addr: portNumber,
		Handler: routes(&app),
	}

	err = srv.ListenAndServe()
	log.Fatal(err)
}