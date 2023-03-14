package main

import (
	"fmt"
	"html/template"
	"net/http"
)

// partials is a package level variable for the template partials
var partials = [...]string{
	"./cmd/web/templates/base.layout.tmpl",
	"./cmd/web/templates/header.partial.tmpl",
	"./cmd/web/templates/footer.partial.tmpl",
}

// render takes an http.Response writer and a template name
// and attempts to parse and execute the template
func render(w http.ResponseWriter, t string) {
	templateSlice := []string{} // creates a new template slice

	templateSlice = append(templateSlice, fmt.Sprintf("./cmd/web/templates/%s", t)) // appends the passed template
	for _, partial := range partials {
		templateSlice = append(templateSlice, partial) // appends all partials
	}

	tmpl, err := template.ParseFiles(templateSlice...) // parses all template contents
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if err := tmpl.Execute(w, nil); err != nil { // executes all templates
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
