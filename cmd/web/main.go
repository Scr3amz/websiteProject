package main

import (
	"log"
	"net/http"
	"os"
	"path/filepath"
)

type application struct {
	infoLog *log.Logger
	errorLog *log.Logger
}

func main() {
	// создаю свой маршрутизатор с ограниченной областью видимости в целях безопасности
	mux := http.NewServeMux()
	port := ":8080"
	fileServer := http.FileServer(restrictedFileSystem{http.Dir("./ui/static/")})

	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	app := application{
		infoLog: infoLog,
		errorLog: errorLog,
	}

	mux.Handle("/static", http.NotFoundHandler())
	mux.Handle("/static/", http.StripPrefix("/static/", fileServer))
	mux.HandleFunc("/", app.index)
	mux.HandleFunc("/about/", app.about_page)

	server := &http.Server{
		Addr: port,
		Handler: mux,
		ErrorLog: errorLog,
	}

	infoLog.Printf("Запуск сервера на порте: %s", port)
	errorLog.Fatal(server.ListenAndServe())

}

type restrictedFileSystem struct {
	fs http.FileSystem
}

func (rfs restrictedFileSystem) Open (path string) (http.File, error) {
	f, err := rfs.fs.Open(path)
	if err != nil {
		return nil, err
	}

	s, err := f.Stat()
	if s.IsDir(){
		index := filepath.Join(path, "index.html")
		if _,err := rfs.fs.Open(index); err != nil {
			closeErr := f.Close()
			if closeErr != nil {
				return nil, closeErr
			}
			return nil, err
		}

	}
	return f, nil
}