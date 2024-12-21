package render

import (
	"bytes"
	//"bytes"
	"html/template"
	"log"
	"net/http"
	"path/filepath"

	"github.com/fspruhs/myGoWebApplication/pkg/config"
	"github.com/fspruhs/myGoWebApplication/pkg/models"
)

var app *config.AppConfig

func NewTemplates(a *config.AppConfig) {
	app = a
}

func AddDefaultData(td *models.TemplateData) *models.TemplateData {
	return td
}

func Template(w http.ResponseWriter, tmpl string, td *models.TemplateData) {
	var tc map[string]*template.Template

	if app.UseCache {
		tc = app.TemplateCache
	} else {
		tc, _ = CreateTemplateCache()
	}

	t, ok := tc[tmpl]
	if !ok {
		log.Fatalln("could not get template from template cache", ok)
	}

	buffer := new(bytes.Buffer)

	td = AddDefaultData(td)

	err := t.Execute(buffer, td)
	if err != nil {
		log.Println("error writing template to buffer", err)
	}

	_, err = buffer.WriteTo(w)
	if err != nil {
		log.Println("error writing template to browser", err)
	}
}

func CreateTemplateCache() (map[string]*template.Template, error) {
	theCache := map[string]*template.Template{}

	pages, err := filepath.Glob("./templates/*-page.gohtml")
	if err != nil {
		return theCache, err
	}

	for _, page := range pages {
		name := filepath.Base(page)
		ts, err := template.New(name).ParseFiles(page)
		if err != nil {
			return theCache, err
		}
		matches, err := filepath.Glob("./templates/*-layout.gohtml")
		if err != nil {
			return theCache, err
		}

		if len(matches) > 0 {
			ts, err = ts.ParseGlob("./templates/*-layout.gohtml")
			if err != nil {
				return theCache, err
			}
		}

		theCache[name] = ts
	}

	return theCache, nil
}
