package mysql

import (
	"database/sql"
	"errors"

	"github.com/Scr3amz/websiteProject/pkg/models"
)

type NoteModel struct {
	DB *sql.DB
}

/*
	Метод для создания новых заметок, принимает название, содержание, и

количество дней, которое заметка будет храниться. Возвращает id заметки,
которую добавил
*/
func (m *NoteModel) Insert(title, content, expires string) (int, error) {
	/*SQL-запрос с использованием плейсхолдеров*/
	sqlCommand := `INSERT INTO notes (title,content,created,expires)
					VALUES (?, ?, UTC_TIMESTAMP(), DATE_ADD(UTC_TIMESTAMP(), INTERVAL ? DAY))`
	resault, err := m.DB.Exec(sqlCommand, title, content, expires)
	if err != nil {
		return -1, err
	}
	id, err := resault.LastInsertId()
	if err != nil {
		return -1, err
	}

	return int(id), nil
}

/*
	Метод, для получения заметки, принимающий её id. Возвращает указатель

на структуру Note, которая заполняется из БД
*/
func (m *NoteModel) Get(id int) (*models.Note, error) {
	sqlCommand := "SELECT * FROM notes WHERE id = ? AND expires > UTC_TIMESTAMP()"
	row := m.DB.QueryRow(sqlCommand, id)
	note := &models.Note{}
	err := row.Scan(&note.ID, &note.Title, &note.Content, &note.Created, &note.Expires)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, models.ErrNoRecord
		} else {
			return nil, err
		}

	}
	return note, nil
}

func (m *NoteModel) Latest() ([]*models.Note, error) {
	return nil, nil
}
