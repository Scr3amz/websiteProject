package models

import (
	"errors"
	"time"
)

var ErrNoRecord = errors.New("models: запись ненайдена")

/* Модель таблицы с заметками */
type Note struct {
	ID      int `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
	Created time.Time `json:"created"`
	Expires time.Time `json:"expires"`
}
