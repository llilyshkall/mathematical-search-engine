package service

import (
	"html/template"
	"net/http"
)

func (s *Service) HomeHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("internal/templates/index.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	tmpl.Execute(w, nil)
}
