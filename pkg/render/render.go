package render

import (
	"fmt"
	"net/http"
	"text/template"
)

func Template(w http.ResponseWriter, tmpl string) {
	parsedTemplate, err := template.ParseFiles("./templates/"+tmpl, "./templates/base-layout.gohtml")
	if err != nil {
		fmt.Println("error parsing template: ", err)
		return
	}

	err = parsedTemplate.Execute(w, nil)
	if err != nil {
		fmt.Println("error parsing template: ", err)
	}
}
