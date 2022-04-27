package main

import (
	"log"
	"net/http"
	"time"

	"github.com/AmanSrivastava2000/bookings/cmd/pkg/config"
	"github.com/AmanSrivastava2000/bookings/cmd/pkg/handlers"
	"github.com/AmanSrivastava2000/bookings/cmd/pkg/render"
	"github.com/alexedwards/scs/v2"
)

const portNo = ":8080"

var app config.AppConfig
var session *scs.SessionManager

func main() {

	//change this to true when in production
	app.InProduction = false

	//setting sessions
	session = scs.New()
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = app.InProduction

	app.Session = session

	//load template cache
	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("cannot create template cache.", err)
	}

	//stroring the tc in templateCache member of app object
	app.TemplateCache = tc
	app.UseCache = false

	//make the repo object by passing the app as to newRepo function in handlers.
	repo := handlers.NewRepo(&app)
	//pass repo object to newHandlers function of handler to initialise the Repo variable there
	handlers.NewHandlers(repo)

	//passing app to render's newTemplates function so that it can be used to get template cache inside it
	render.NewTemplates(&app)

	//setting up the routes:
	srv := &http.Server{
		Addr:    portNo,
		Handler: routes(&app),
	}

	err = srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}

}
