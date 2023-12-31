package handlers

import "net/http"

// Маршрутизатор с ограниченной областью видимости в целях безопасности
func (app *Application) Router() *http.ServeMux {
	mux := http.NewServeMux()
	fileServer := http.FileServer(http.Dir("./ui/static/"))

	// Обработчики статических и html- файлов

	mux.HandleFunc("/", app.index)
	mux.HandleFunc("/notes/", app.notes_page)
	mux.HandleFunc("/notes/show/", app.show_note)
	mux.HandleFunc("/notes/create/", app.create_note)

	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	return mux
}
