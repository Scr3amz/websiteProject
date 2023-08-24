package main

import (
	"log"
	"net/http"
	"path/filepath"
)

func main() {
	// создаю свой маршрутизатор с ограниченной областью видимости в целях безопасности
	mux := http.NewServeMux()
	fileServer := http.FileServer(restrictedFileSystem{http.Dir("./static/")})
	mux.Handle("/static", http.NotFoundHandler())
	mux.Handle("/static/", http.StripPrefix("/static/", fileServer))
	mux.HandleFunc("/", index)
	mux.HandleFunc("/about/", about_page)

	log.Fatal(http.ListenAndServe(":8080", mux))

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