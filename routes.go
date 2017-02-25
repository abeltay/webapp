package main

import (
	"html/template"
	"net/http"
)

var templateDir = "template/"

var templates = template.Must(template.ParseFiles(
	templateDir+"index.tmpl",
	templateDir+"1.tmpl",
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

func One(w http.ResponseWriter, r *http.Request) {
	p := &Page{
		Title: "1",
	}
	renderTemplate(w, "1", p)
}

func Two(w http.ResponseWriter, r *http.Request) {
	p := &Page{
		Title: "shortlist",
	}
	renderTemplate(w, "shortlist", p)
}

func renderTemplate(w http.ResponseWriter, tmpl string, data *Page) {
	err := templates.ExecuteTemplate(w, tmpl, data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
