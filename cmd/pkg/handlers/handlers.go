package handlers

import (
	"net/http"

	"github.com/AmanSrivastava2000/bookings/cmd/pkg/config"
	"github.com/AmanSrivastava2000/bookings/cmd/pkg/models"
	"github.com/AmanSrivastava2000/bookings/cmd/pkg/render"
)

//Template data hold data sent from handlers to templates.

//Repository is the repository type
type Repository struct{
	App *config.AppConfig
}

// Repo is repository used by handlers
var Repo *Repository

//creates and a new Repository
func NewRepo(a *config.AppConfig)(*Repository){
	return &Repository{
		App: a,
	}
}

//sets the repository for handlers.
func NewHandlers(r *Repository){
	Repo = r
}

//home page handler
func (m *Repository)Home(w http.ResponseWriter, r *http.Request) {
	//everytime someone hits the home page, extract its ip_address and store in session member
	remoteIP := r.RemoteAddr
	m.App.Session.Put(r.Context(), "remote_ip", remoteIP)

	render.RenderTemplate(w, "home.page.tmpl", &models.TemplateData{})
}

//about page handler
func (m *Repository)About(w http.ResponseWriter, r *http.Request) {
	//make data to pass
	//get the ip_address:
	remoteIP := m.App.Session.GetString(r.Context(), "remote_ip")

	//put the data into stringMap
	stringMap := make(map[string]string)
	stringMap["test"] = "hello this was the data passed from handlers."
	stringMap["remote_ip"] = remoteIP

	render.RenderTemplate(w, "about.page.tmpl", &models.TemplateData{
		StringMap: stringMap,
	})
}
