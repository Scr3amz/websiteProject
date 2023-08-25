package main

import (
	"html/template"
	"net/http"
)

func (app *application) index(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	files := []string{
		"./ui/html/index.html",
		"./ui/html/header.html",
		"./ui/html/footer.html",
	}
	t, err := template.ParseFiles(files...)
	if err != nil {
		app.errorLog.Println(err.Error())
		http.Error(w, "Internal Server Error", 500)
		return
	}
	err = t.ExecuteTemplate(w, "index", nil)
	if err != nil {
		app.errorLog.Println(err.Error())
		http.Error(w, "Internal Server Error", 500)
		return
	}
}

func (app *application)about_page(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/about/" {
		http.NotFound(w, r)
		return
	}
	files := []string{
		"./ui/html/about.html",
		"./ui/html/header.html",
		"./ui/html/footer.html",
	}
	t, err := template.ParseFiles(files...)
	if err != nil {
		app.errorLog.Println(err.Error())
		http.Error(w, "Internal Server Error", 500)
		return
	}
	err = t.ExecuteTemplate(w, "about", nil)
	if err != nil {
		app.errorLog.Println(err.Error())
		http.Error(w, "Internal Server Error", 500)
		return
	}

}
