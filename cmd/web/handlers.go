package main

import (
	"errors"
	"fmt"
	"html/template"
	"net/http"
	"strconv"

	"github.com/Scr3amz/websiteProject/pkg/models"
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
func (app *application) about_page(w http.ResponseWriter, r *http.Request) {
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

/*Обработчик страницы с заметками*/
func (app *application) notes_page(w http.ResponseWriter, r *http.Request) {
	// Получение параметра id из ссылки
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil || id < 1 {
		http.NotFound(w, r)
		return
	}

	note, err := app.notes.Get(id)
	if err != nil {
		if errors.Is(err, models.ErrNoRecord) {
			app.notFound(w)
			return
		} else {
			app.serverError(w, err)
		}
	}
	fmt.Fprintf(w, "%v", note)
	/*
		files := []string{
			"./ui/html/notes.html",
			"./ui/html/header.html",
			"./ui/html/footer.html",
		}
		t, err := template.ParseFiles(files...)
		if err != nil {
			app.serverError(w, err)
			return
		}
		err = t.ExecuteTemplate(w, "notes", note)
		if err != nil {
			app.serverError(w, err)
			return
		}
	*/
}

/* Обработчик для создания заметки*/
func (app *application) create_note(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		app.clientError(w, http.StatusMethodNotAllowed)
		return
	}

	title := "test note"
	content := "test content"
	expires := "2"

	id, err := app.notes.Insert(title, content, expires)
	if err != nil {
		app.serverError(w, err)
		return
	}

	http.Redirect(w, r, fmt.Sprintf("/notes?id=%d", id), http.StatusSeeOther)

}
