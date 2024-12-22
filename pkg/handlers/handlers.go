package handlers

import (
	"net/http"

	"github.com/fspruhs/myGoWebApplication/pkg/config"
	"github.com/fspruhs/myGoWebApplication/pkg/models"
	"github.com/fspruhs/myGoWebApplication/pkg/render"
)

type Repository struct {
	App *config.AppConfig
}

var Repo *Repository

func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{
		App: a,
	}
}

func NewHandlers(r *Repository) {
	Repo = r
}

func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {
	remoteIP := r.RemoteAddr
	m.App.Session.Put(r.Context(), "remote_ip", remoteIP)
	render.Template(w, "home-page.gohtml", &models.TemplateData{
		StringMap: map[string]string{"morty": "smith"},
	})
}

func (m *Repository) About(w http.ResponseWriter, r *http.Request) {
	remoteIP := m.App.Session.GetString(r.Context(), "remote_ip")
	render.Template(w, "about-page.gohtml", &models.TemplateData{
		StringMap: map[string]string{"remote_ip": remoteIP},
	})
}
