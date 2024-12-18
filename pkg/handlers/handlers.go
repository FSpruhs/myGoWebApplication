package handlers

import (
	"net/http"

	"github.com/fspruhs/myGoWebApplication/pkg/render"
)

func Home(w http.ResponseWriter, r *http.Request) {
	render.Template(w, "home-page.gohtml")
}

func About(w http.ResponseWriter, r *http.Request) {
	render.Template(w, "about-page.gohtml")
}
