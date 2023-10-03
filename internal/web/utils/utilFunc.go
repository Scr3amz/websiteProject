package utils

import (
	"database/sql"

	"github.com/Scr3amz/websiteProject/pkg/models"
)

/* Структура-оболочка для передачи сразу нескольких динамических элементов*/
type templateData struct {
	Note *models.Note
	Notes []*models.Note
}


func CreateData() *templateData {
	 return &templateData{}
}

/* Функция для открытия базы данных и проверки соединения с ней */
func OpenDB(dsn string) (*sql.DB, error) {
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}
	if err = db.Ping(); err != nil {
		return nil, err
	}
	return db, nil
}