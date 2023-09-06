package models

import (
	"errors"
	"time"
)

var ErrNoRecord = errors.New("models: запись ненайдена")

/* Модель таблицы с заметками */
type Note struct {
	ID      int
	Title   string
	Content string
	Created time.Time
	Expires time.Time
}
