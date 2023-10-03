package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/Scr3amz/websiteProject/internal/web/config"
	"github.com/Scr3amz/websiteProject/internal/web/handlers"
	"github.com/Scr3amz/websiteProject/pkg/models/mysql"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	config, err := config.LoadConfig()
	if err!= nil {
        log.Fatal("Failed to load configuration\n",err)
    }

	/* Логгеры для более удобной и настраевоемой обработки ошибок и
	выводе информации в терминал */
	InfoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	ErrorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	/* Подключение к пулу соединений с базой данных
	"web:8803@/mvol_website?parseTime=true" */
	db, err := openDB( config.DBUser + ":" + config.DBPass + "@/" + config.DBName + "?" + config.DriverParams)
	if err != nil {
		ErrorLog.Fatal(err)
	}
	defer db.Close()
	

	app := handlers.Application{
		InfoLog:  InfoLog,
		ErrorLog: ErrorLog,
		Notes:    &mysql.NoteModel{DB: db},
	}

	/* Отдельный объект для сервера чтобы указать в качестве поля ErrorLog свой логгер, а
	в качестве обработчика функцию, создающую маршрутизатор*/
	server := &http.Server{
		Addr:     config.Port,
		Handler:  app.Router(),
		ErrorLog: ErrorLog,
	}

	InfoLog.Printf("Запуск сервера по адресу: 127.0.0.1%s", config.Port)
	ErrorLog.Fatal(server.ListenAndServe())

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
