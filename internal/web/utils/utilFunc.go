package utils

import (
	"database/sql"
	"encoding/json"
	"net/http"

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

func DecodeJSON(r *http.Request, v interface{}) error {
	defer r.Body.Close()
	decoder := json.NewDecoder(r.Body)
    return decoder.Decode(v)
}