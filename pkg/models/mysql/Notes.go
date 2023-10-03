package mysql

import (
	"database/sql"
	"errors"

	"github.com/Scr3amz/websiteProject/pkg/models"
)

type NoteModel struct {
	DB *sql.DB
}

/* Метод для создания новых заметок, принимает название, содержание, и
количество дней, которое заметка будет храниться. Возвращает id заметки,
которую добавил */
func (m *NoteModel) Insert(title, content, expires string) (int, error) {
	/*SQL-запрос с использованием плейсхолдеров*/
	sqlCommand := `INSERT INTO notes (title,content,created,expires)
	VALUES (?, ?, LOCALTIMESTAMP(), DATE_ADD(LOCALTIMESTAMP(), INTERVAL ? DAY))`
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

/* Метод, для получения заметки, принимающий её id. Возвращает указатель
на структуру Note, которая заполняется из БД */
func (m *NoteModel) Get(id int) (*models.Note, error) {
	sqlCommand := "SELECT * FROM notes WHERE id = ? AND expires > LOCALTIMESTAMP()"
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

/* Метод, для получения последних 10 заметок. Возвращает указатель на срез
 структур Note, которые заполняюся из БД */
func (m *NoteModel) Latest() ([]*models.Note, error) {
	sqlCommand := "SELECT * FROM notes WHERE expires > LOCALTIMESTAMP() ORDER BY created DESC LIMIT 10"

	rows, err := m.DB.Query(sqlCommand)
	if err!=nil {
		return nil, err
	}
	defer rows.Close()
	notes := []*models.Note{}

	for rows.Next() {
		note := &models.Note{}
		err = rows.Scan(&note.ID, &note.Title, &note.Content, &note.Created, &note.Expires )
		if err!=nil {
			return nil, err
		}
		notes = append(notes, note)
	}

	err = rows.Err() 
	if err!=nil {
		return nil, err
	}
	
	return notes, nil
}
