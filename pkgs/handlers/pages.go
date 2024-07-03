package handlers

import (
	"github.com/sungatishe/pkgs/config"
	"github.com/sungatishe/pkgs/models"
	"github.com/sungatishe/pkgs/renders"
	"net/http"
)

var Repo *Repository

type Repository struct {
	App *config.AppConfig
}

func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{
		App: a,
	}
}

func NewHandlers(r *Repository) {
	Repo = r
}

func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {
	stringMap := make(map[string]string)
	stringMap["test"] = "Hello!"

	remoteIP := m.App.Session.GetString(r.Context(), "remoteIP")
	stringMap["remoteIP"] = remoteIP

	renders.RenderTemplate(w, "home.page.tmpl", &models.TemplateData{
		StringMap: stringMap,
	})
}

func (m *Repository) About(w http.ResponseWriter, r *http.Request) {
	remoteIP := r.RemoteAddr
	m.App.Session.Put(r.Context(), "remoteIP", remoteIP)
	renders.RenderTemplate(w, "about.page.tmpl", &models.TemplateData{})
}
