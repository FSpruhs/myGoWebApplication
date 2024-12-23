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
	render.Template(w, "about-page.gohtml", &models.TemplateData{})
}

func (m *Repository) Contact(w http.ResponseWriter, r *http.Request) {
	render.Template(w, "contact-page.gohtml", &models.TemplateData{})
}

func (m *Repository) Eremite(w http.ResponseWriter, r *http.Request) {
	render.Template(w, "eremite-page.gohtml", &models.TemplateData{})
}

func (m *Repository) Couple(w http.ResponseWriter, r *http.Request) {
	render.Template(w, "couple-page.gohtml", &models.TemplateData{})
}

func (m *Repository) Family(w http.ResponseWriter, r *http.Request) {
	render.Template(w, "family-page.gohtml", &models.TemplateData{})
}

func (m *Repository) Reservation(w http.ResponseWriter, r *http.Request) {
	render.Template(w, "check-availability-page.gohtml", &models.TemplateData{})
}

func (m *Repository) MakeReservation(w http.ResponseWriter, r *http.Request) {
	render.Template(w, "make-reservation-page.gohtml", &models.TemplateData{})
}
