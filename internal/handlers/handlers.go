package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/fspruhs/myGoWebApplication/internal/config"
	"github.com/fspruhs/myGoWebApplication/internal/forms"
	"github.com/fspruhs/myGoWebApplication/internal/models"
	"github.com/fspruhs/myGoWebApplication/internal/render"
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
	render.Template(w, r, "home-page.gohtml", &models.TemplateData{
		StringMap: map[string]string{"morty": "smith"},
	})
}

func (m *Repository) About(w http.ResponseWriter, r *http.Request) {
	render.Template(w, r, "about-page.gohtml", &models.TemplateData{})
}

func (m *Repository) Contact(w http.ResponseWriter, r *http.Request) {
	render.Template(w, r, "contact-page.gohtml", &models.TemplateData{})
}

func (m *Repository) Eremite(w http.ResponseWriter, r *http.Request) {
	render.Template(w, r, "eremite-page.gohtml", &models.TemplateData{})
}

func (m *Repository) Couple(w http.ResponseWriter, r *http.Request) {
	render.Template(w, r, "couple-page.gohtml", &models.TemplateData{})
}

func (m *Repository) Family(w http.ResponseWriter, r *http.Request) {
	render.Template(w, r, "family-page.gohtml", &models.TemplateData{})
}

func (m *Repository) Reservation(w http.ResponseWriter, r *http.Request) {
	render.Template(w, r, "check-availability-page.gohtml", &models.TemplateData{})
}

type jsonResponse struct {
	OK      bool   `json:"ok"`
	Message string `json:"message"`
}

func (m *Repository) ReservationJSON(w http.ResponseWriter, r *http.Request) {
	resp := jsonResponse{
		OK:      true,
		Message: "Available!",
	}

	output, err := json.MarshalIndent(resp, "", "    ")
	if err != nil {
		log.Println(err)
	}

	w.Header().Set("Content-Type", "application/json")
	_, _ = w.Write(output)
}

func (m *Repository) PostReservation(w http.ResponseWriter, r *http.Request) {
	start := r.Form.Get("startingDate")
	end := r.Form.Get("endingDate")
	w.Write([]byte(fmt.Sprintf("Arrival is %s , depature ist %s", start, end)))
}

func (m *Repository) MakeReservation(w http.ResponseWriter, r *http.Request) {
	var emptyReservation models.Reservation

	data := make(map[string]interface{})
	data["reservation"] = emptyReservation

	render.Template(w, r, "make-reservation-page.gohtml", &models.TemplateData{
		Data: data,
		Form: forms.New(nil),
	})
}

func (m *Repository) PostMakeReservation(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		log.Println(err)
		return
	}

	reservation := models.Reservation{
		Name:  r.Form.Get("full_name"),
		Email: r.Form.Get("email"),
		Phone: r.Form.Get("phone"),
	}

	form := forms.New(r.PostForm)

	form.Required("full_name", "email")
	form.MinLength("full_name", 2)
	form.IsEmail("email")

	if !form.Valid() {
		data := make(map[string]interface{})
		data["reservation"] = reservation
		render.Template(w, r, "make-reservation-page.gohtml", &models.TemplateData{
			Form: form,
			Data: data,
		})
		return
	}
}
