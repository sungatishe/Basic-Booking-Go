package main

import (
	"github.com/sungatishe/pkgs/config"
	"github.com/sungatishe/pkgs/handlers"
	"github.com/sungatishe/pkgs/renders"
	"log"
	"net/http"
	"time"

	"github.com/alexedwards/scs/v2"
)

const portNumber = ":8080"

var app config.AppConfig
var session *scs.SessionManager

func main() {

	app.InProd = false

	session = scs.New()
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = app.InProd

	app.Session = session

	tc, err := renders.CreateCacheTemplate()

	if err != nil {
		log.Fatal("cannot create tmpl")
	}

	app.TemplateCheck = tc
	app.UseCache = false

	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)

	renders.NewTemplates(&app)

	srv := &http.Server{
		Addr:    portNumber,
		Handler: route(),
	}
	err = srv.ListenAndServe()
	log.Fatal(err)
}
