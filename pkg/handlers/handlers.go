package handlers

import (
	"net/http"

	"github.com/conrad3rd/bookings/pkg/config"
	"github.com/conrad3rd/bookings/pkg/models"
	"github.com/conrad3rd/bookings/pkg/render"
)

// Repo the repository used by th handlers
var Repo *Reposetory

// Reposetory is the repository type
type Reposetory struct {
	App *config.AppConfig
}

// NewRepo creates a new repository
func NewRepo(a *config.AppConfig) *Reposetory {
	return &Reposetory{
		App: a,
	}
}

// NewHanlders sets the repository for the handlers
func NewHandlers(r *Reposetory) {
	Repo = r
}

// Home is the home page handler
func (m *Reposetory) Home(w http.ResponseWriter, r *http.Request) {
	remoteIP := r.RemoteAddr
	m.App.Session.Put(r.Context(), "remote_ip", remoteIP)

	render.RenderTemplate(w, "home.page.tmpl", &models.TemplateData{})
}

// About is the about page handler
func (m *Reposetory) About(w http.ResponseWriter, r *http.Request) {
	// perform some logic
	stringMap := make(map[string]string)
	stringMap["test"] = "Hello, again"

	remoteIP := m.App.Session.GetString(r.Context(), "remote_ip")
	stringMap["remote_ip"] = remoteIP

	// send the data to the template
	render.RenderTemplate(w, "about.page.tmpl", &models.TemplateData{
		StringMap: stringMap,
	})
}
