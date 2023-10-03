package handlers

import (
	"errors"
	"fmt"
	"html/template"
	"net/http"
	"strconv"

	"github.com/Scr3amz/websiteProject/internal/web/utils"
	"github.com/Scr3amz/websiteProject/pkg/models"
)



// Обработчик главной страницы
func (app *Application) index(w http.ResponseWriter, r *http.Request) {
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

/*Обработчик страницы с заметками*/
func (app *Application) notes_page(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/notes/" {
		app.notFound(w)
		return
	}

	notes, err := app.Notes.Latest()
	if err != nil {
		app.serverError(w, err)
		return
	}

	data := utils.CreateData()
	data.Notes = notes
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
	err = t.ExecuteTemplate(w, "notes", data)
	if err != nil {
		app.serverError(w, err)
		return
	}

}

/* Метод, показывающий заметку по её id из БД*/
func (app *Application) show_note(w http.ResponseWriter, r *http.Request) {
	// Получение параметра id из ссылки
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil || id < 1 {
		http.NotFound(w, r)
		return
	}

	note, err := app.Notes.Get(id)
	if err != nil {
		if errors.Is(err, models.ErrNoRecord) {
			app.notFound(w)
			return
		} else {
			app.serverError(w, err)
		}
	}
	files := []string{
		"./ui/html/show.html",
		"./ui/html/header.html",
		"./ui/html/footer.html",
	}
	data := utils.CreateData()
	data.Note = note
	t, err := template.ParseFiles(files...)
	if err != nil {
		app.serverError(w, err)
		return
	}
	err = t.ExecuteTemplate(w, "show", data)
	if err != nil {
		app.serverError(w, err)
		return
	}
}

/* Обработчик для создания заметки. Ждёт POST-запрос в json-формате
с указанием названия заметки (title) и её содержания (content)*/
func (app *Application) create_note(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		app.clientError(w, http.StatusMethodNotAllowed)
		return
	}

	var note models.Note
	err := utils.DecodeJSON(r, &note)
	if err!= nil {
		fmt.Println(err)
        app.clientError(w, http.StatusBadRequest)
        return
    }

	id, err := app.Notes.Insert(note.Title, note.Content, "1")
	if err != nil {
		app.serverError(w, err)
		return
	}

	http.Redirect(w, r, fmt.Sprintf("/notes?id=%d", id), http.StatusSeeOther)

}
