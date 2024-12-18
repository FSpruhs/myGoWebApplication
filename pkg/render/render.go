package render

import (
	"bytes"
	"html/template"
	"log"
	"net/http"
	"path/filepath"
)

func Template(w http.ResponseWriter, tmpl string) {

	t, ok := tc[tmpl]
	if !ok {
		log.Fatalln("could not get template from template cache", err)
	}

	buffer := new(bytes.Buffer)

	err = t.Execute(buffer, nil)
	if err != nil {
		log.Println("error writing template to browser", err)
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
