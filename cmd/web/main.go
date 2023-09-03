package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"
	"path/filepath"

	_ "github.com/go-sql-driver/mysql"
)

/*
	Структура для передачи логгеров в соседние файлы, и

для общей инкапсуляции приложения
*/
type application struct {
	infoLog  *log.Logger
	errorLog *log.Logger
}

func main() {
	port := ":8080"

	/* Логгеры для более удобной и настраевоемой обработки ошибок и
	выводе информации в терминал */
	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	/* Подключение к базе данных */
	db, err := openDB("web:8803@/mvol_website?parseTime=true")
	if err != nil {
		errorLog.Fatal(err)
	}

	defer db.Close()

	app := application{
		infoLog:  infoLog,
		errorLog: errorLog,
	}

	/* Отдельный объект для сервера чтобы указать в качестве поля ErrorLog свой логгер, а
	в качестве обработчика функцию, создающую маршрутизатор*/
	server := &http.Server{
		Addr:     port,
		Handler:  app.router(),
		ErrorLog: errorLog,
	}

	infoLog.Printf("Запуск сервера на порте: %s", port)
	errorLog.Fatal(server.ListenAndServe())

}

// Ограничиваю доступ к просмотру файлов, если в директории нет index-файла

type restrictedFileSystem struct {
	fs http.FileSystem
}

func (rfs restrictedFileSystem) Open(path string) (http.File, error) {
	f, err := rfs.fs.Open(path)
	if err != nil {
		return nil, err
	}

	s, err := f.Stat()
	if s.IsDir() {
		index := filepath.Join(path, "index.html")
		if _, err := rfs.fs.Open(index); err != nil {
			closeErr := f.Close()
			if closeErr != nil {
				return nil, closeErr
			}
			return nil, err
		}

	}
	return f, nil
}

/* Функция для открытия базы данных и проверки соединения с ней */

func openDB(dsn string) (*sql.DB, error) {
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}
	if err = db.Ping(); err != nil {
		return nil, err
	}
	return db, nil
}
