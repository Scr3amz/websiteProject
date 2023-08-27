package main

import (
	"html/template"
	"net/http"
)

// Обработчик главной страницы
func (app *application) index(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		app.notFound(w)
		return
	}
	files := []string{
		"./ui/html/index.html",
		"./ui/html/header.html",
		"./ui/html/footer.html",
	}
	t, err := template.ParseFiles(files...)
	if err != nil {
		app.serverError(w, err)
		return
	}
	err = t.ExecuteTemplate(w, "index", nil)
	if err != nil {
		app.serverError(w, err)
		return
	}
}

// Обработчик страницы "обо мне"
func (app *application)about_page(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/about/" {
		app.notFound(w)
		return
	}
	files := []string{
		"./ui/html/about.html",
		"./ui/html/header.html",
		"./ui/html/footer.html",
	}
	t, err := template.ParseFiles(files...)
	if err != nil {
		app.serverError(w, err)
		return
	}
	err = t.ExecuteTemplate(w, "about", nil)
	if err != nil {
		app.serverError(w, err)
		return
	}

}
