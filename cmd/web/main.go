package main

import (
	"log"
	"net/http"
	"os"

	"github.com/Scr3amz/websiteProject/internal/web/config"
	"github.com/Scr3amz/websiteProject/internal/web/handlers"
	"github.com/Scr3amz/websiteProject/internal/web/utils"
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

	/* Подключение к пулу соединений с базой данных */
	db, err := utils.OpenDB( config.DBUser + ":" + config.DBPass + "@/" + config.DBName + "?" + config.DriverParams)
	if err != nil {
		ErrorLog.Fatal(err)
	}
	defer db.Close()
	

	app := handlers.Application{
		InfoLog:  InfoLog,
		ErrorLog: ErrorLog,
		Notes:    &mysql.NoteModel{DB: db},
	}

	/* Отдельный объект для сервера чтобы подключить свой логгер, а
	в качестве обработчика функцию, создающую маршрутизатор */
	server := &http.Server{
		Addr:     config.Port,
		Handler:  app.Router(),
		ErrorLog: ErrorLog,
	}

	InfoLog.Printf("Запуск сервера по адресу: 127.0.0.1%s", config.Port)
	ErrorLog.Fatal(server.ListenAndServe())

}
