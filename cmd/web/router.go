package main

import "net/http"

// Маршрутизатор с ограниченной областью видимости в целях безопасности
func (app *application) router() *http.ServeMux {
	mux := http.NewServeMux()
	fileServer := http.FileServer(restrictedFileSystem{http.Dir("./ui/static/")})

	// Обработчики статических и html- файлов
	mux.Handle("/static", http.NotFoundHandler())
	mux.Handle("/static/", http.StripPrefix("/static/", fileServer))
	mux.HandleFunc("/", app.index)
	mux.HandleFunc("/about/", app.about_page)
	mux.HandleFunc("/notes/", app.notes_page)
	mux.HandleFunc("/notes/create/", app.create_note)

	return mux
}
