package main

import (
	"html/template"
	"net/http"
)

var templateDir = "template/"

var templates = template.Must(template.ParseFiles(
	templateDir+"index.tmpl",
	templateDir+"budget.tmpl",
	templateDir+"shortlist.tmpl",
	templateDir+"header.tmpl",
	templateDir+"footer.tmpl"))

type Page struct {
	Title string
}

func Index(w http.ResponseWriter, r *http.Request) {
	p := &Page{
		Title: "Home",
	}
	renderTemplate(w, "index", p)
}

func Budget(w http.ResponseWriter, r *http.Request) {
	p := Page{
		Title: "Budget Calculation",
	}
	renderTemplate(w, "budget", p)
}

func Shortlist(w http.ResponseWriter, r *http.Request) {
	p := struct {
		Title string
		Areas []string
		MRT   []string
		Paid  bool
	}{
		Title: "Shortlist",
		Areas: PlanningArea,
		MRT:   MRT,
	}
	renderTemplate(w, "shortlist", p)
}

func renderTemplate(w http.ResponseWriter, tmpl string, data interface{}) {
	err := templates.ExecuteTemplate(w, tmpl, data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
