
package controllers

import (
	"net/http"
	"html/template"
)

type HomeController struct {
	BaseController
}

func (h HomeController) Index(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("resources/views/index.html")
	if err != nil {
		http.Error(w, "Error loading template", http.StatusInternalServerError)
		return
	}
	tmpl.Execute(w, nil)
}
