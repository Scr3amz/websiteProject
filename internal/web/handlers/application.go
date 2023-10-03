package handlers

import (
	"log"

	"github.com/Scr3amz/websiteProject/pkg/models/mysql"
)

/*
	Структура для передачи логгеров в соседние файлы, и

для общей инкапсуляции приложения
*/
type Application struct {
	InfoLog  *log.Logger
	ErrorLog *log.Logger
	Notes    *mysql.NoteModel
}